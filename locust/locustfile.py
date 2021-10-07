from locust import HttpUser, task


class HelloWorldUser(HttpUser):
    @task
    def get_apis(self):
        self.client.get("/node/sha256")
        self.client.get("/go/sha256")

    @task
    def post_apis(self):
        self.client.post("/node/sha256", {"Input": "AAAAAAAAAAAAAAAAAAA"})
        self.client.post("/go/sha256", {"Input": "AAAAAAAAAAAAAAAAAAA"})
