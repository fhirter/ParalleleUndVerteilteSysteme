# Übung Virtualisierung

## Lernziele

- Die Studierenden kennen verschiedene Virtualisierungstechniken und deren Anwendung.
- Die Studierenden können Applikationen containerisieren und container orchestrieren.

## Vorbereitung

Lade und installiere [Docker Desktop](https://www.docker.com/products/docker-desktop/).

## Aufgabenstellung

Halte alle Schritte in einem [Markdown-Dokument](https://www.markdownguide.org/) fest.

Stelle sicher, dass alle ausgeführten Befehle vollständig nachvollziehbar sind!

**Schicke das Dokument nach dem Unterricht an den Dozenten.**

### Applikation

Bearbeite folgende Kapitel des[ Go Docker Tutorials](https://docs.docker.com/language/golang/).

1. Overview
2. Build Images
3. Run Containers
4. Develop your app

Nutze die [beiliegende Applikation](main.go).

Überspringe allenfalls das Kapitel "Local database and containers" um für die nächste Aufgabe mehr Zeit zu haben.

### Monitoring

Erweitere die Applikation mit Monitoring.

1. Füge diese beiden Applikationen zum docker compose file hinzu:
    1. [Prometheus](https://prometheus.io/docs/prometheus/latest/installation/)
    2. [Grafana](https://grafana.com/docs/grafana/latest/setup-grafana/installation/docker/)
2. Erweitere den Applikations-Code um einen einfachen Metrics
   Endpoint: [examples/simple/main.go](https://github.com/prometheus/client_golang/blob/main/examples/simple/main.go)
3. [Konfiguriere Prometheus](https://prometheus.io/docs/prometheus/latest/configuration/configuration/) so, dass der
   metrics Endpoint abgefragt wird.
4. Erstelle ein Dashboard in Grafana für die geladenen Metrics.