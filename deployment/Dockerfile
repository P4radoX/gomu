# MIT License

# Copyright (c) 2021 P4radoX

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

##
# You can change base image, however, alpine is the default choice to reduce image size.
# The micro-service binary must be compiled before building the image. Only the executable
# binary will be copied to image layer. The binary arguments can be passed at container execution
# due to container-entrypoint.sh script which run the micro-service as PID 1 in the container environment.
##

FROM alpine:3.12.4

# Labels
LABEL version="v1.0.0"
LABEL name="Micro-service name"
LABEL description="Micro-service description"

# Environment vars
ENV GOMU_APP_BIN="micro-service-binary"
ENV GOMU_PORT="8080"
ENV GOMU_VOLUME_DATA="/var/lib/micro-service_name/data"
ENV GOMU_VOLUME_SSL="/var/lib/micro-service_name/ssl"
ENV GOMU_URL_ENDPOINT="/micro-service-name/v1/"

# Copy entrypoint script & binary
COPY ./container-entrypoint ./
COPY ./${GOMU_APP_BIN} /bin

# Make directories
RUN mkdir -p ${GOMU_VOLUME_DATA}
RUN mkdir -p ${GOMU_VOLUME_SSL}

# Volumes
VOLUME [ ${GOMU_VOLUME_DATA} ]
VOLUME [ ${GOMU_VOLUME_SSL} ]

# Healthcheck
HEALTHCHECK --interval=1m --timeout=5s --start-period=20s --retries=3 CMD [ "curl -X GET http://localhost:${GOMU_PORT}${GOMU_URL_ENDPOINT}/health || exit 1" ]

# Expose
EXPOSE 8080

# Entrypoint
ENTRYPOINT [ "./container-entrypoint.sh" ]

# Run container, modify by the golang compiled binary name
CMD [ ${GOMU_APP_BIN} ]