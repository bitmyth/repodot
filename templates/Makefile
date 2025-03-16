# 定义变量
PROJECT_NAME = myproject
GO_VERSION = 1.20
PYTHON_VERSION = 3.9
NODE_VERSION = 16

# 默认目标
.DEFAULT_GOAL := help

# 帮助信息
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help      显示帮助信息"
	@echo "  build     构建项目"
	@echo "  test      运行测试"
	@echo "  clean     清理构建文件"
	@echo "  run       运行项目"
	@echo "  docker    构建 Docker 镜像"
	@echo "  deploy    部署项目"

# 构建项目
build:
	@echo "Building $(PROJECT_NAME)..."
	# 这里可以添加具体的构建命令
	go build -o bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)

# 运行测试
test:
	@echo "Running tests..."
	# 这里可以添加具体的测试命令
	go test -v ./...

# 清理构建文件
clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -rf dist/
	rm -rf coverage/

# 运行项目
run:
	@echo "Running $(PROJECT_NAME)..."
	./bin/$(PROJECT_NAME)

# 构建 Docker 镜像
docker:
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME):latest .

# 部署项目
deploy:
	@echo "Deploying $(PROJECT_NAME)..."
	# 这里可以添加具体的部署命令
	kubectl apply -f k8s/deployment.yaml

# 安装依赖
install:
	@echo "Installing dependencies..."
	# 这里可以添加具体的依赖安装命令
	go mod download
	npm install
	pip install -r requirements.txt

# 格式化代码
format:
	@echo "Formatting code..."
	# 这里可以添加具体的格式化命令
	gofmt -w .
	black .
	prettier --write .

# 检查代码风格
lint:
	@echo "Linting code..."
	# 这里可以添加具体的代码检查命令
	golint ./...
	flake8 .
	eslint .

# 生成文档
docs:
	@echo "Generating documentation..."
	# 这里可以添加具体的文档生成命令
	swagger generate spec -o ./swagger.yaml
	mkdocs build

# 发布版本
release:
	@echo "Releasing $(PROJECT_NAME)..."
	# 这里可以添加具体的发布命令
	git tag v1.0.0
	git push origin v1.0.0

# 运行开发服务器
dev:
	@echo "Starting development server..."
	# 这里可以添加具体的开发服务器命令
	nodemon --exec go run main.go