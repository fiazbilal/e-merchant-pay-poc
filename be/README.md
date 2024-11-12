### Setup
#### Environment Configuration:
Create a .env file in the api/cmd folder, following the same structure as the env.sh file. This file should contain all necessary environment variables for the application to run.

#### Running the Application:
To start the application, execute the following script from the root directory:

```bash
./devrun.sh
```

### Docker Setup

#### Build the Docker Image
From the root directory, build the Docker image by running:

```bash
docker build -t payment-processor .
```

#### Run the Docker Container
Once the image is built, run a container from it:

```bash
docker run -p 1786:1786 --env-file cmd/.env payment-processor
```

- This command maps port 1786 in the container to port 1786 on your local machine.
- The --env-file cmd/.env option loads environment variables from the .env file into the container.

### Testing
To test the application, you can use the following curl command to create a payment:

```bash
curl -X POST http://localhost:1786/payment/create \
     -H "Content-Type: application/json" \
     -d '{
           "transaction_id": "6fdda189-9425-4a9b-b9ee-f734776e64cf",
           "amount": 1000,
           "currency": "EUR"
         }'
```
Replace "transaction_id", "amount", and "currency" with your desired values.