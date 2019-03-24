import grpc

from protos.python.weather_pb2 import WeatherRequest
from protos.python import weather_pb2_grpc

from google.protobuf.json_format import MessageToDict


class WeatherClient(object):

    def __init__(self):
        channel = grpc.insecure_channel('weather_service:50051')
        self.stub = weather_pb2_grpc.WeatherServiceStub(channel)

    def get(self, request, meta=None):

        request_obj = request
        if type(request) == dict:
            request_obj = WeatherRequest(**request)

        return MessageToDict(self.stub.Get(request_obj))
