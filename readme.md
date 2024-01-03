# Go project boilerplate

## Description
This is a boilerplate for a Go project. By UXCentury
Used best practices and patterns for a Go project.

## Usage

Make sure to give it execute permissions: 
```sh
chmod +x scripts/build.sh # build the project
chmod +x scripts/run-dev.sh # run the project
```

### Notes

For ci/cd use .gitlab-ci.yml

For local docker build use Dockerfile


#### Go Version - 1.21.5

#### Database - Postgres

#### Utils
- [Zap](https://github.com/uber-go/zap) logger
- [Viper](https://github.com/spf13/viper) config
- Retry (custom retry logic)