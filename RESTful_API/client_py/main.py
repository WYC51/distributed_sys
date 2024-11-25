import requests
import argparse
import json


API_URL = "localhost"

def CLI_args():

    parser = argparse.ArgumentParser(description="RESTful API client")

    parser.add_argument("--method", type = str, default = "GET", help = "CRUD operation")

    parser.add_argument("--id", type = int, default = -1, help = "ID of the data")

    parser.add_argument("--port", type = int, default = 8080, help = "Port number")

    return parser.parse_args()

def get_data(id, port):
    if id == -1:
        url = f"http://{API_URL}:{port}/show_demo_data"
        response = requests.get(url).json()
        return response
    
    else:
        url = f"http://{API_URL}:{port}/show_demo_data/{id}"
        response = requests.get(url).json()
        return response

def post_data(port):

    url = f"http://{API_URL}:{port}/show_demo_data"

    data = {
        "name": "test500",
        "id": 500,
        "mail": "500"
    }
    headers = {'Content-Type' : 'application/json'}
    response = requests.post(url, data = json.dumps(data), headers=headers).json()

    return response

def put_data(id, port):

    url = f"http://{API_URL}:{port}/show_demo_data/{id}"

    data = {
        "name": "test500",
        "id": 500,
        "mail": "500@update.com"
    }

    headers = {'Content-Type' : 'application/json'}

    response = requests.put(url, data = json.dumps(data), headers=headers).json()

    return response

def delete_data(id, port):

    url = f"http://{API_URL}:{port}/show_demo_data/{id}"

    response = requests.delete(url).json()

    return response

def main(args):
    if args.method == "GET":
        response = get_data(args.id, args.port)
        print(response)

    elif args.method == "POST":
        response = post_data(args.port)
        print(response)

    elif args.method == "PUT":
        response = put_data(args.id, args.port)
        print(response)

    elif args.method == "DELETE":
        response = delete_data(args.id, args.port)
        print(response)

if __name__ == "__main__":
    args = CLI_args()
    main(args)
    
