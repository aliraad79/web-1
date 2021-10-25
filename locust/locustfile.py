from locust import HttpUser, task
import json


class HelloWorldUser(HttpUser):
    def on_start(self):
        response = self.client.post("/node/sha256", {"Input": "test input",})
        self.sha256 = response.json().get('SHA')
    
    @task
    def get_apis(self):
        self.client.get(f"/node/sha256?sha={self.sha256}")
        self.client.get(f"/go/sha256?sha={self.sha256}")

    @task
    def post_apis(self):
        self.client.post("/node/sha256", {"Input": "AAAAAAAAAAAAAAAAAAA"})
        self.client.post("/go/sha256", json.dumps({"Input": "AAAAAAAAAAAAAAAAAAA"}))
