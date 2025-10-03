# ================================
# Stage 1: Build the Vue + Vite project
# ================================

FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json if exists
COPY package*.json ./

# Install all dependencies including devDependencies
RUN npm install

# Copy all source code
COPY . .

# Build the project
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

# Expose port used by Vite preview
EXPOSE 4173

# Run preview server
CMD ["npm", "run", "preview"]
