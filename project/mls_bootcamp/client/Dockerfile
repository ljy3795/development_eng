# (ref. https://mherman.org/blog/dockerizing-a-react-app/)
# docker build -t mls-bootcamp-react .

#base image
FROM node:alpine3.10

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
RUN npm install --silent
RUN npm install react-scripts -g --silent

# start app
CMD ["yarn", "start"]


# docker run -d -p 3000:3000 --rm -v ${PWD}:/app -v /app/node_modules --name go-mls-bootcamp-react mls-bootcamp-react:latest