# (ref. https://mherman.org/blog/dockerizing-a-react-app/)
# docker build --no-cache -t mls-bootcamp-react .

#base image
FROM node:alpine3.10

# set working directory
WORKDIR /app

COPY ./src ./src
COPY ./public ./public/
RUN ls -la /app/

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
# package.json에 있는 거 설치 (nodejs만 있는 곳에서)
RUN npm install --silent
RUN npm install react-scripts -g --silent

# start app
CMD ["yarn", "start"]