FROM node:lts-alpine
ENV NODE_ENV=production
WORKDIR /usr/src/app
COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]
RUN npm install --production --silent && mv node_modules ../
COPY . .
EXPOSE 3000
EXPOSE 8226
EXPOSE 8228
EXPOSE 7003
EXPOSE 43300
RUN chown -R node /usr/src/app
USER node
CMD ["node", "index.js"]
