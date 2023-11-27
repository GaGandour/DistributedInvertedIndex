import sys

def create_worker(worker_id: int) -> str:
    return f"""  worker-{worker_id}:
    image: dii
    container_name: worker-{worker_id}
    command: ["./dii", "-type", "worker", "-addr", "worker-{worker_id}", "-port", "500{worker_id:02d}", "-master", "dii-master:5000"]
    ports:
      - "500{worker_id:02d}:500{worker_id:02d}"
"""

def create_master() -> str:
    return """  dii-master:
    image: dii
    container_name: dii-master
    command: ["./dii", "-type", "master", "-addr", "dii-master"]
    tty: true
    stdin_open: true
    ports:
      - "5000:5000"
"""

def create_docker_compose(num_workers: int) -> str:
    workers_description = "\n".join([create_worker(i) for i in range(1, num_workers+1)])
    return f"""version: '3'
services:
{create_master()}
{workers_description}"""

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python write_docker_compose.py <num_workers>")
        sys.exit(1)
    print(create_docker_compose(int(sys.argv[1])))