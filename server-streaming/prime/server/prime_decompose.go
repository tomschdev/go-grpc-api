package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/prime/proto"
)

func (s *Server) PrimeDecomposition(in *pb.PrimeRequest, stream pb.PrimeService_PrimeDecompositionServer) error {

	log.Printf("PrimeDecomposition was invoked with: %v\n", in)
	var k uint32 = 2
	for N := in.Number; N > 1; {
		if N%k == 0 {
			log.Printf("Identified prime factor: %s\n", k)
			stream.Send(&pb.PrimeResponse{Prime: k})
			N = N / k
		} else {
			k = k + 1
		}
	}
	log.Printf("PrimeDecomposition completed")
	return nil
}
