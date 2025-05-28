PROJECT_ID := hinaproject-460723
SERVICE_NAME := english-practice
REGION := us-central1
PORT := 8080

deploy:
	@echo "Deploying ${SERVICE_NAME} to Cloud Run..."
	gcloud run deploy ${SERVICE_NAME} \
		--project=${PROJECT_ID} \
		--region=${REGION} \
		--source=. \
		--allow-unauthenticated \
		--port=${PORT}

url:
	gcloud run services describe ${SERVICE_NAME} --project=${PROJECT_ID} --region=${REGION} --format="value(status.url)"

