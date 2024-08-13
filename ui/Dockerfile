# ui/Dockerfile
# Use official Node.js image as base image
FROM node:18

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application
RUN npm run build

# Expose port 8080
EXPOSE 8080

# Start the Vue.js application
CMD ["npm", "run", "serve"]