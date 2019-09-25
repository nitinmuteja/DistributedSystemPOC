package org.dist.queue

import java.util.concurrent.atomic.AtomicLong

class Replica(val brokerId: Int,
              val partition: Partition,
              time: Time = SystemTime,
              initialHighWatermarkValue: Long = 0L,
              val log: Option[Log] = None) extends Logging {
  //only defined in local replica
  private[this] var highWatermarkValue: AtomicLong = new AtomicLong(initialHighWatermarkValue)
  // only used for remote replica; logEndOffsetValue for local replica is kept in log
  private[this] var logEndOffsetValue = new AtomicLong(ReplicaManager.UnknownLogEndOffset)
  private[this] var logEndOffsetUpdateTimeMsValue: AtomicLong = new AtomicLong(time.milliseconds)
  val topic = partition.topic
  val partitionId = partition.partitionId

  def logEndOffset_=(newLogEndOffset: Long) {
    if (!isLocal) {
      logEndOffsetValue.set(newLogEndOffset)
      logEndOffsetUpdateTimeMsValue.set(time.milliseconds)
      trace("Setting log end offset for replica %d for partition [%s,%d] to %d"
        .format(brokerId, topic, partitionId, logEndOffsetValue.get()))
    } else
      throw new KafkaException("Shouldn't set logEndOffset for replica %d partition [%s,%d] since it's local"
        .format(brokerId, topic, partitionId))

  }

  def logEndOffset = {
    if (isLocal)
      log.get.logEndOffset
    else
      logEndOffsetValue.get()
  }

  def isLocal: Boolean = {
    log match {
      case Some(l) => true
      case None => false
    }
  }

  def logEndOffsetUpdateTimeMs = logEndOffsetUpdateTimeMsValue.get()

  def highWatermark_=(newHighWatermark: Long) {
    if (isLocal) {
      trace("Setting hw for replica %d partition [%s,%d] on broker %d to %d"
        .format(brokerId, topic, partitionId, brokerId, newHighWatermark))
      highWatermarkValue.set(newHighWatermark)
    } else
      throw new KafkaException("Unable to set highwatermark for replica %d partition [%s,%d] since it's not local"
        .format(brokerId, topic, partitionId))
  }

  def highWatermark = {
    if (isLocal)
      highWatermarkValue.get()
    else
      throw new KafkaException("Unable to get highwatermark for replica %d partition [%s,%d] since it's not local"
        .format(brokerId, topic, partitionId))
  }

  override def equals(that: Any): Boolean = {
    if(!(that.isInstanceOf[Replica]))
      return false
    val other = that.asInstanceOf[Replica]
    if(topic.equals(other.topic) && brokerId == other.brokerId && partition.equals(other.partition))
      return true
    false
  }

  override def hashCode(): Int = {
    31 + topic.hashCode() + 17*brokerId + partition.hashCode()
  }


  override def toString(): String = {
    val replicaString = new StringBuilder
    replicaString.append("ReplicaId: " + brokerId)
    replicaString.append("; Topic: " + topic)
    replicaString.append("; Partition: " + partition.partitionId)
    replicaString.append("; isLocal: " + isLocal)
    if(isLocal) replicaString.append("; Highwatermark: " + highWatermark)
    replicaString.toString()
  }
}
