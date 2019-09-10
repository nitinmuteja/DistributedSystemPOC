/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.dist.dbgossip;

import java.io.IOException;
import java.util.Collections;
import java.util.List;

/**
 * This is the first message that gets sent out as a start of the Gossip protocol in a
 * round.
 */
public class GossipDigestSyn
{
    final String clusterId;
    final String partioner;
    final List<GossipDigest> gDigests;

    public GossipDigestSyn(String clusterId, String partioner, List<GossipDigest> gDigests)
    {
        this.clusterId = clusterId;
        this.partioner = partioner;
        this.gDigests = gDigests;
    }

    List<GossipDigest> getGossipDigests()
    {
        return gDigests;
    }
}

class GossipDigestSerializationHelper
{
    static void serialize(List<GossipDigest> gDigestList, int version) throws IOException
    {

    }

    static List<GossipDigest> deserialize(/*DataInputPlus in,*/ int version) throws IOException
    {
        return Collections.EMPTY_LIST;
    }

    static int serializedSize(List<GossipDigest> digests, int version)
    {

        return 0;
    }
}

class GossipDigestSynSerializer
{
    public void serialize(GossipDigestSyn gDigestSynMessage, /*DataOutputPlus out,*/ int version) throws IOException
    {

    }

    public GossipDigestSyn deserialize(/*DataInputPlus in,*/ int version) throws IOException
    {
            return null;
    }

    public long serializedSize(GossipDigestSyn syn, int version)
    {
        return 0;
    }
}

