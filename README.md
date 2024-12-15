# AQUATECH IoT Platform
> Smart Aquaculture Monitoring System

## Overview

AquaTech is a real-time IoT monitoring platform for fish farms, providing automated data collection, analysis, and marketplace integration. Built with Go and modern web technologies, it helps farmers prevent losses using data analysis from IoT devices and access premium markets.

## Project Structure


aqua-tech/
├── Arduino/           # IoT sensor firmware
├── cmd/              # Application entry points
│   └── main.go       # Server initialization
├── db/               # Database migrations & schemas
├── static/           # Static assets & styles
├── templates/        # HTML templates
└── Dockerfile        # Container configuration


## Features

### Core Functionality
- Real-time sensor data collection and analysis
- Automated monitoring & alerts
- Farm management dashboard
- Marketplace integration
- Community features

### Data Collection
- Temperature monitoring
- Oxygen Concentration
- Water level sensing (ultrasonic)
- Data storage

## Tech Stack

- *Backend*: Go
- *Frontend*: HTML, CSS, JavaScript
- *IoT*: Arduino-based sensors
- *Containerization*: Docker

## Quick Start

1. *Prerequisites*
bash
go 
arduino
docker


2. *Setup*
bash
# Clone repository
git clone https://github.com/Aaron-Ochieng/aqua-tech.git
cd aqua-tech

# Build and run
go mod download
go run cmd/main.go


3. *Docker Deployment*
bash
docker build -t aqua-tech .
docker run -p 8080:8080 aqua-tech


## Development

### Server
bash
# Run development server
go run cmd/main.go

# Build binary
go build -o aquatech cmd/main.go


### Arduino Setup
1. Flash the sensor firmware from Arduino/ directory
2. Configure sensor parameters
3. Connect to the main system

## API Endpoints


GET  /              - Dashboard
GET  /farms         - Farm Management
GET  /market        - Marketplace
GET  /community     - Community Features
POST /api/data      - Sensor Data Ingestion


## Configuration

Key configuration files:
- Dockerfile: Container settings
- Arduino/config.h: Sensor configurations
- Environment variables (see .env.example)

## Data Structure

go
type Data struct {
    Temp           float64
    Humidity       float64
    UltraSonicData float64
}


## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## Development Guidelines

- Follow Go best practices
- Document API changes
- Test sensor integrations
- Update documentation

## Troubleshooting

Common issues:
1. Sensor connectivity: Check Arduino configuration
2. Template errors: Verify template paths
3. Data flow: Monitor sensor data pipeline

## Support

- Technical Issues: Create GitHub issue
- Hardware Support: Check Arduino documentation
- General Queries: Contact development team

## License

Proprietary - All rights reserved

---

*Note*: This README is maintained by the development team. For latest updates, check the repository.

## Team

- Tomlee Abila `https://github.com/Tomlee-abila`
- Valentine Omollo `https://github.com/vomolo`
- Emmanuel Barsulai `https://github.com/Barsu5489`
- Abraham Kingoo `https://github.com/abrakingoo`
- Aaron Ochieng `https://github.com/Aaron-Ochieng`

