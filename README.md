<div align="center">
  
# Trapwire

![Go version](https://img.shields.io/github/go-mod/go-version/NotAProton/trapwire?filename=app%2Fgo.mod&logo=Go)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/notaproton/trapwire/docker-publish.yml?logo=Github&label=Build)
![GitHub Tag](https://img.shields.io/github/v/tag/notaproton/trapwire?label=Latest%20version)
![License](https://img.shields.io/github/license/notaproton/trapwire)

</div>
<br>

 🔗🚨 **Trapwire** is a lightweight, Go-powered link shortener with built-in surveillance for cyber security and analytics. Easily deployable as a Docker container, Trapwire helps you monitor and analyze URL access for threat detection and user insights.

## Features
- 🔗 **URL Shortening**: Create shortened links for easy sharing.
- 📸 **Surveillance**: Track link access including IP, time visited, and user agents.
- 🛡️ **Cybersecurity**: Monitor for suspicious behavior and malicious access attempts.
- 📊 **Analytics**: Gather real-time insights into traffic and user behavior.
- 🐳 **Dockerized**: Simple setup and deployment with Docker.

## Installation

### Docker compose
Clone the repo, edit the docker-compose file to change credentials and run it.

```bash
git clone https://github.com/NotAProton/trapwire.git
cd trapwire
nano docker-compose.yml
docker-compose up --d
```

### Docker
To run Trapwire using Docker with your own database and reverse proxy setup:

```bash
docker pull ghcr.io/notaproton/trapwire:latest
docker run -p 3000:3000 -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e AUTH=your_password -e TOKEN=hello ghcr.io/notaproton/trapwire
```
