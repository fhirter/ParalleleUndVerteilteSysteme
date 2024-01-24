# Übung Prozesse

Halte dein Vorgehen, Überlegungen und Erkenntnisse in einem Dokument fest.
Schicke dieses am Ende der Unterrichtseinheit dem Dozenten.

Organisiert euch selbstständig in der Klasse, wie ihr die Aufgabe Lösen wollt.
Stellt sicher, dass alle verstehen, was gemacht wurde.

## Lernziele

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.
- Die Studierenden kennen Kommunikations- und Koordinationstechniken zwischen Prozessen und Systemen

## Aufgabenstellung

Halte die Erkenntnisse in einem [Markdown-Dokument](https://www.markdownguide.org/) fest.

Stelle sicher, dass deine Untersuchungen vollständig nachvollziehbar sind!

**Schicke das Dokument nach dem Unterricht an den Dozenten.**

### Systemanalyse

Verbinde dich mit den separat kommunizierten Zugangsdaten mit dem Linux-System.

Untersuche anhand der laufenden Prozesse, wozu das System dient.
Identifiziere die wichtigsten Prozesse.

Untersuche diese identifizierten Prozesse mit `ps` und `htop` indem du folgende Fragen beantwortest.
Beschränke dich dabei nicht nur auf die reine beantwortung der Frage, sondern erkläre die Frage an sich sowie deinen
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
2. Die Rückgabewerte werden mittels Channels übergeben, da normale Funktionsrückgabewerte mit asynchroner Ausführung
   nicht funktionieren. Erkläre wieso!
3. Führe die drei Funktionen in Goroutines aus. Miss die Zeit, die die Berechnung benötigt.
4. Erstelle eine Statistik, die aufzeigt, wann in diesem Beispiel die Multiprocessing Lösung schneller ist.

### Scheduling

The folgenden Prozesse werden von einem präemptiven, round-robin Scheduling-Algorithmus verwaltet:

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

Wenn ein Task von einem anderen mit höherer Priorität unterbrochen wird, wird der unterbrochene Task ans Ende der
Warteschlange gestellt.

1. Zeichne ein Gantt Diagramm für die Prozessabfolge.
2. Wie hoch ist die CPU-Auslastung?
3. Wie gross ist die Wartezeit für jeden Prozess?

a. Show the scheduling order of the processes using a Gantt chart.
b. What is the turnaround time for each process?
c. What is the waiting time for each process?
d. What is the CPU utilization rate?
