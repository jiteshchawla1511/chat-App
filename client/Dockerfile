
# Use the official Node.js image as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the local package files to the container's workspace
COPY . .

# Install dependencies
RUN npm install

# Build the React application
RUN npm run build

# Expose the port on which the React app will run (if necessary)
# EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
