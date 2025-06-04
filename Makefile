PROJECT_ID := hinaproject-460723
SERVICE_NAME := english-practice
REGION := us-central1
PORT := 8080
IMAGE_URL := ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/${IMAGE}

local:
	docker build -t img:local .
	docker run -p 8081:8080 img:local

build:
	docker build -t us-central1-docker.pkg.dev/hinaproject-460723/english-practice-repo/english-practice:latest .

deploy:
	@echo "Deploying ${SERVICE_NAME} to Cloud Run..."
	gcloud builds submit --tag us-central1-docker.pkg.dev/hinaproject-460723/english-practice-repo/english-practice:staging
	gcloud run deploy english-practice --image us-central1-docker.pkg.dev/hinaproject-460723/english-practice-repo/english-practice:staging  --region ${REGION}
url:
	gcloud run services describe ${SERVICE_NAME} --project=${PROJECT_ID} --region=${REGION} --format="value(status.url)"

