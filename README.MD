@author Michael Dara

Instructions on how to build and run the project on a local Mac machine

1. install docker and start the docker machine

2. log in to the docker machine

3.  Open a terminal and create the below Environment Variables 
    export AWS_REGION = <aws region name>
    export AWS_ACCESS_KEY_ID = <your access key id>
    export AWS_SECRET_ACCESS_KEY = <your secret key>

4. Clone the below git repo and cd to s3query folder
    git clone xxx

5. Build the docker image
    docker build -t mdara/s3query . 

6. Run the docker image on the docker container with different test cases.

    test-1: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query
    
    test-1 result: 
        input flag cannot be blank


    test-2: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query -input s3://maf-sample-data/1k.ndjson.gz
    
    test-2 result: 
        all 1000 records returned from the s3 file 1k.ndjson.gz


    test-3: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query -input s3://maf-sample-data/1k.ndjson.gz -with-id 310942977652378514
    
    test-3 result: 
        records matching the with-id flag are returned from the s3 file 1k.ndjson.gz


    test-4: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query -input s3://maf-sample-data/1k.ndjson.gz -with-word=scribner
    
    test-4 result: 
        records matching the with-word flag are returned from the s3 file 1k.ndjson.gz


     test-5: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query -input s3://maf-sample-data/1k.ndjson.gz -from-time=2000-01-01T00:00:00Z -to-time=2001-01-01T00:00:00Z
    
    test-5 result: 
        records matching the from-time and to-time flags are returned from the s3 file 1k.ndjson.gz


     test-6: 
        docker run --rm -e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY mdara/s3query -input s3://maf-sample-data/1k.ndjson.gz -from-time=2000-01-01T00:00:00Z -to-time=2015-01-01T00:00:00Z --with-word=scribner
    
    test-6 result: 
        records matching the from-time,to-time and with-word flags are returned from the s3 file 1k.ndjson.gz      
