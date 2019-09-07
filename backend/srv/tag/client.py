import grpc

from proto import tag_pb2_grpc, tag_pb2

with grpc.insecure_channel('localhost:8500') as channel:
    stub = tag_pb2_grpc.TagStub(channel)
    response = stub.GetTags(tag_pb2.TagRequest(description="舒肤佳沐浴露"))
    for tag in response.tags:
        print(tag)
