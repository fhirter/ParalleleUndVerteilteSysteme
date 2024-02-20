# Übung Prozesse

## Lernziele

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.
- Die Studierenden kennen Kommunikations- und Koordinationstechniken zwischen Prozessen und Systemen

## Vorgehen

Halte dich an das folgende, wissenschaftliche Vorgehen:

1. Fragestellung formulieren
2. Recherche
3. Vorgehen festlegen
4. Untersuchungen durchführen
5. Frage beantworten
6. Ausblick

Stelle sicher, dass du alle Fragen mit belastbaren Daten beantworten kannst.

Halte die Erkenntnisse in einem [Markdown-Dokument](https://www.markdownguide.org/) fest.

**Schicke das Dokument nach dem Unterricht an den Dozenten.**

## Aufgaben

### Systemanalyse

Verbinde dich mit den separat kommunizierten Zugangsdaten mit dem Linux-System.

Untersuche anhand der laufenden Prozesse, wozu das System dient.
Identifiziere die wichtigsten Prozesse.

Untersuche diese identifizierten Prozesse mit `ps` und `htop` indem du folgende Fragen beantwortest.
Beschränke dich dabei nicht nur auf die reine Beantwortung der Frage, sondern erkläre die Frage an sich sowie deinen
Lösungsprozess.

- Welche Prozesse nutzen wie viel Speicher?
- Welche Prozesse nutzen wie viel CPU?
- Welche Prozesse haben wie viel CPU Zeit verwendet?
- Welchen Ausführungsstatus haben die Prozesse?
- Welche Scheduling Priorität haben die Prozesse?
- Wie viele Prozesse starten die einzelnen Applikationen?
- Welche Prozesse nutzen mehrere Threads?

### Multiprocessing

Entwickle ein Programm, das verschiedene statistische Werte für eine Liste von Zahlen berechnet.
Das Programm soll eine Liste von zufälligen Zahlen generieren und davon den Durchschnitt, den maximalen und den
minimalen Wert berechnen.

Zum Beispiel, nehmen wir an, das Programm generiert die Zahlenreihe `90 81 78 95 79 72 85`.
Das Programm berechnet daraus:

- Der Durchschnittswert ist 82
- Der Minimalwert ist 72
- Der Maximalwert ist 95

Verwende für die Lösung [Go](https://go.dev/), dieses bietet sehr einfaches Multiprocessing.

Nutze die beliegende Vorlage aus Ausgangslage.

1. Führe die drei Berechnungen in einem Prozess aus. Miss die Zeit, die die Berechnung benötigt.
2. Führe die drei Funktionen in Goroutines aus. Miss die Zeit, die die Berechnung benötigt.
3. Erstelle eine Statistik, die die gemessenen Zeiten in Abhängigkeit der Länge der Zahlenreihe darstellt.

### Kommunikation

Optimiere die Berechnung für den Durchschnittswert aus der vorherigen Übung.

Nutze für die Kommunikation zwischen den Goroutines und dem Hauptprogramm buffered channels; Einen Channel für die
Daten, einen für die Resultate.

Buffered channels erstellst du mit `messages := make(chan string, 2)`. Der zweite Parameter von make ist die Grösse des
Buffers, d.h. die Anzahl Elemente, die im Channel gespeichert werden können.

Führe für jeden Optimierungsschritt Messungen durch, damit du die Optimierung mit Zahlen belegen kannst.

1. Recherchiere, wie Go Channels funktionieren und vergleiche mit der Beschreibung im Kapitel 3.6 (Silberschatz, 2019).
2. Schreibe das ganze Array in den Daten-Channel. Starte eine Goroutine, die die Daten aus dem Channel liest, den
   Mittelwert berechnet und das Resultat in den Resultat-Channel schreibt.
3. Unterteile die Daten in mehrere kleinere Arrays, die du in den Daten-Channel schreibst. Ändere die Goroutine so, dass
   sie läuft, bis keine Daten mehr im Daten-Channel vorhanden sind.
4. Starte nun mehrere Goroutines, die die Daten verarbeiten.

### Scheduling

Löse folgende Aufgabe für verschiedene Scheduling Algorithmen:

- [Priority based, preemptive, round robin](https://en.wikipedia.org/wiki/Fixed-priority_pre-emptive_scheduling)
- [Completely Fair Scheduler](https://en.wikipedia.org/wiki/Completely_Fair_Scheduler)
- [Multilevel Feedback Queue](https://de.wikipedia.org/wiki/Multilevel_Feedback_Queue)

| Bezeichner | Priorität | CPU Zeit | Ready Zeit |
|------------|-----------|----------|------------|
| P1         | 40        | 20       | 0          |
| P2         | 30        | 25       | 25         |
| P3         | 30        | 25       | 30         |
| P4         | 35        | 15       | 60         |
| P5         | 5         | 10       | 100        |
| P6         | 10        | 10       | 105        |

Jeder Prozess hat eine Priorität, wobei eine höhere Zahl eine höhere Priorität bedeutet.
Zusätzlich zu den aufgeführten Prozessen existiert ein Idle-Task, der ausgeführt wird, wenn keine anderen Prozesse
verfügbar sind.
Er besitzt Priorität 0.

Der Scheduler wird alle 10 Zeiteinheiten ausgeführt.

1. Zeichne ein Gantt Diagramm für die Prozessabfolge.
2. Wie hoch ist die CPU-Auslastung?
3. Wie gross ist die Wartezeit für jeden Prozess?
