# build stage (frontend)
FROM node:lts-alpine as build-stage

# Update aptitude with new repo
RUN apt-get update

# Install software 
RUN apt-get install -y git

RUN git clone https://github.com/NexSabre/mikropoker.git .

WORKDIR /app
COPY package*.json ./

RUN yarn -g install
COPY . .
RUN yarn build

# build stage (backend & package)
FROM golang:1.15.7-buster



# Install Python stuff
RUN pip install cython
ADD ./requirements.txt /tmp/requirements.txt

RUN pip install --no-cache-dir -q -r /tmp/requirements.txt


# Copy files
ADD . /opt/packet_server/
RUN mkdir -p /opt/packet_server/tmp
RUN mkdir -p /opt/packet_server/static

WORKDIR /opt/packet_server
COPY --from=build-stage /app/dist/static static
COPY --from=build-stage /app/dist/index.html static

# Set ENV's
ARG PH_REV
ENV PH_REVISION=${PH_REV}

ARG PH_VER
ENV PH_VERSION=${PH_VER}

# $PORT is set by Heroku
CMD uvicorn --host 0.0.0.0 --port $PORT ph.main:app