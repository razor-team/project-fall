FROM ubuntu:latest

RUN apt-get update

RUN apt-get install wget gnupg

RUN wget -qO - http://packages.lunarg.com/lunarg-signing-key-pub.asc | apt-key add -
RUN wget -qO /etc/apt/sources.list.d/lunarg.list \
            http://packages.lunarg.com/vulkan/1.2.131/lunarg-vulkan-1.2.131-bionic.list

RUN apt-get intall -y cmake ninja-build libgtest-dev libglm-dev libxcb-dri3-0 \
                    libxcb-present0 libpciaccess0 libpng-dev libxcb-keysyms1-dev \
                    libxcb-dri3-dev libx11-dev libmirclient-dev \
                    libwayland-dev libxrandr-dev libglfw3-dev git python \
                    mesa-vulkan-drivers lunarg-vulkan-sdk
