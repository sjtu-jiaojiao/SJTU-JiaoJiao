import time
import grpc
import consul
import json
import getopt
import sys
from concurrent import futures

from proto import tag_pb2_grpc, tag_pb2

from config import path_hyper_parameters
from text_preprocess import PreprocessTextMulti, load_json
from graph import TextCNNGraph as Graph
import numpy as np

registry_address = "127.0.0.1"
register_ttl = 20
register_interval = 60


class TagHandler(tag_pb2_grpc.TagServicer):
    def get_tags(self, request, context):
        hyper_parameters = load_json(path_hyper_parameters)
        pt = PreprocessTextMulti()
        # 模式初始化和加载
        graph = Graph(hyper_parameters)
        graph.load_model()
        ra_ed = graph.word_embedding

        # str to token
        ques_embed = ra_ed.sentence2idx(request["description"])
        if hyper_parameters['embedding_type'] == 'bert':
            x_val_1 = np.array([ques_embed[0]])
            x_val_2 = np.array([ques_embed[1]])
            x_val = [x_val_1, x_val_2]
        else:
            x_val = ques_embed
        pred = graph.predict(x_val)

        pre = pt.prereocess_idx(pred[0])
        ls_nulti = []
        for ls in pre[0]:
            if ls[1] >= 0.5:
                ls_nulti.append(ls)

        json_response = tag_pb2.TagResponse()
        json_response.rst_string = json.dumps(ls_nulti)
        return json_response


def register(server_name, ip, port):
    c = consul.Consul()  # 连接consul 服务器，默认是127.0.0.1，可用host参数指定host
    print(f"开始注册服务{server_name}")
    check = consul.Check.tcp(ip, port, register_interval, register_ttl)  # 健康检查的ip，端口，检查时间
    c.agent.service.register(server_name, f"{server_name}-{ip}-{port}",
                             address=ip, port=port, check=check)  # 注册服务部分
    print(f"注册服务{server_name}成功")


def unregister(server_name, ip, port):
    c = consul.Consul()
    print(f"开始退出服务{server_name}")
    c.agent.service.deregister(f'{server_name}-{ip}-{port}')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    tag_pb2_grpc.add_TagServicer_to_server(TagHandler(), server)
    server.add_insecure_port('[::]:{}'.format(8500))
    register("tag_server", registry_address, 8500)
    server.start()
    try:
        while True:
            time.sleep(186400)
    except KeyboardInterrupt:
        unregister("tag_server", registry_address, 8500)
        server.stop(0)


try:
    opts, args = getopt.getopt(sys.argv[1:], "h", ["help", "registry_address=", "register_ttl=",
                                                   "register_interval="])  # sys.argv[1:] 过滤掉第一个参数(它是脚本名称，不是参数的一部分)
    for cmd, arg in opts:  # 使用一个循环，每次从opts中取出一个两元组，赋给两个变量。cmd保存选项参数，arg为附加参数。接着对取出的选项参数进行处理。
        if cmd in ("-h", "--help"):
            print("help info")
            sys.exit()
        elif cmd == "--registry_address":
            registry_address = arg
        elif cmd == "--register_ttl":
            register_ttl = arg
        elif cmd == "--register_interval":
            register_interval = arg
    serve()

except getopt.GetoptError:
    print("argv error,please input")
