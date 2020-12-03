"""This is a simple python demo client"""
import grpc
from addresspb import address_pb2
from addresspb import address_pb2_grpc


channel = grpc.insecure_channel('localhost:50051')

stub = address_pb2_grpc.AddressParseServiceStub(channel)

address = address_pb2.APRequest(address="12303 airport way ste 250 broomfield co 80021")

response = stub.AddressParse(address)

print(response)
