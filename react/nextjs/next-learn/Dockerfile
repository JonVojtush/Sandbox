# Use the official Node.js image from the Docker Hub
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (if available) to the working directory
COPY package*.json ./

# Install Node.js dependencies
RUN npm install

# Copy the rest of your application code to the working directory
COPY . .

# Expose the port your app runs on
EXPOSE 3000

# Command to run your application
CMD ["node", "app.js"]
