#aws dynamodb list-tables --endpoint-url=http://localhost:4566

# dynamo
aws --endpoint-url=http://localhost:4566 dynamodb create-table \
    --table-name connections \
    --attribute-definitions \
        AttributeName=user,AttributeType=S \
    --key-schema AttributeName=user,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST
#aws --endpoint-url=http://localhost:4566 dynamodb create-table \
#    --table-name messages \
#    --attribute-definitions \
#        AttributeName=uuid,AttributeType=S \
#    --key-schema \
#        AttributeName=uuid,KeyType=HASH \
#    --billing-mode PAY_PER_REQUEST

# sqs
aws --endpoint-url=http://localhost:4566 sqs create-queue \
--queue-name messages

#aws --endpoint-url=http://localhost:4566 dynamodb put-item \
#    --table-name connections  \
#    --item \
#    '{"user":{"S":"red"}}'
#
#aws --endpoint-url=http://localhost:4566 dynamodb scan \
#    --table-name connections
#
#
#aws --endpoint-url=http://localhost:4566 dynamodb delete-item \
#    --table-name messages  \
#    --key \
#    '{"uuid":{"S":"5c3798de-f259-4cd3-96fe-01c070050945"}}'