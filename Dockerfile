FROM ubuntu:latest

RUN apt update && apt install -y \
    iproute2 \
    iptables \
    net-tools \
    curl \
    traceroute \
    tcpdump \
    dnsutils \
    vim \
    sudo \
    && apt clean

# Set timezone to Pacific Standard Time (PST)
ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Expose a terminal shell as entry point
CMD ["tail", "-f", "/dev/null"]