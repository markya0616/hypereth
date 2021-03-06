// Copyright 2018 AMIS Technologies
// This file is part of the hypereth library.
//
// The hypereth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The hypereth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the hypereth library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// AddPeer connects to the given nodeURL.
func (ec *Client) AddPeer(ctx context.Context, nodeURL string) error {
	var r bool
	return ec.c.CallContext(ctx, &r, "admin_addPeer", nodeURL)
}

// BatchAddPeer performs batch add remote peers.
func (ec *Client) BatchAddPeer(ctx context.Context, urls []string) error {
	if len(urls) == 0 {
		return nil
	}
	// Construct batch requests
	method := "admin_addPeer"
	reqs := make([]rpc.BatchElem, len(urls))
	for i, url := range urls {
		reqs[i] = rpc.BatchElem{
			Method: method,
			Args:   []interface{}{url},
		}
	}
	// Batch calls
	err := ec.c.BatchCallContext(ctx, reqs)
	if err != nil {
		return err
	}
	// Ensure all requests are ok
	for _, req := range reqs {
		if req.Error != nil {
			return err
		}
	}
	return nil
}

// AdminPeers returns the number of connected peers.
func (ec *Client) AdminPeers(ctx context.Context) ([]*p2p.PeerInfo, error) {
	var r []*p2p.PeerInfo
	err := ec.c.CallContext(ctx, &r, "admin_peers")
	if err != nil {
		return nil, err
	}
	return r, err
}

// NodeInfo gathers and returns a collection of metadata known about the host.
func (ec *Client) NodeInfo(ctx context.Context) (*p2p.PeerInfo, error) {
	var r *p2p.PeerInfo
	err := ec.c.CallContext(ctx, &r, "admin_nodeInfo")
	if err != nil {
		return nil, err
	}
	return r, err
}
