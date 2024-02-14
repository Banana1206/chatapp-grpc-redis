const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

// Load file .proto và gRPC package
const PORT = 8082;
const PROTO_FILE = './proto/chat.proto';

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE));

const protoDescriptor = grpc.loadPackageDefinition(packageDef);
const chatApp = protoDescriptor.chatapp;

// Tạo một client gRPC mới
const client = new chatApp.ChatApp(`0.0.0.0:50069`, grpc.credentials.createInsecure());

// Gửi yêu cầu tới server
client.Add({ number: 10, anotherNumber: 29 }, (error, response) => {
    if (!error) {
        console.log('Sum:', response.sum.toString());
    } else {
        console.error('Error:', error);
    }
});
