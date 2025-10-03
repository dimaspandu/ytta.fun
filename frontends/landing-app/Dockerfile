# ================================
# Stage 1: Build the Astro project
# ================================

FROM node:20-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy package.json and package-lock.json if exists
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all source code
COPY . .

# Build the Astro project
RUN npm run build

# ================================
# Stage 2: Runtime image
# ================================

FROM node:20-alpine

WORKDIR /app

# Copy built files from builder stage
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package*.json ./

# Install only production dependencies
RUN npm install --production

# Expose default port used by Astro preview
EXPOSE 4173

# Run the built project
CMD ["npm", "run", "preview"]
