# MouseMover

Simple application to move the mouse cursor randomly every 30 seconds.

## Prerequisites

- Docker
- Docker Compose
- An X11 server for Windows (e.g., [VcXsrv](https://sourceforge.net/projects/vcxsrv/) or [Xming](https://sourceforge.net/projects/xming/))

## Installation

1. **Install Docker and Docker Compose**:
   - Download and install [Docker Desktop](https://www.docker.com/products/docker-desktop) for Windows.

2. **Install an X11 server**:
   - Download and install [VcXsrv](https://sourceforge.net/projects/vcxsrv/) or [Xming](https://sourceforge.net/projects/xming/).
   - Launch the X11 server.

3. **Configure X11 display on Windows**:
   - Open a terminal on your Windows machine and run the following command to allow X11 connections:
     ```sh
     xhost +local:docker
     ```

## Usage

1. **Clone the repository**:
   ```sh
   git clone https://github.com/your-username/mousemover.git
   cd mousemover