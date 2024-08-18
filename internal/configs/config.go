package configs

import "time"

const ServerAddress = "localhost:8080"
const FileName = "../../data/sample_registrations"
const SleepTimeForSimulation = 15 * time.Second
const MaxResends = 3
const ConnectionTimeout = 10 * time.Second
