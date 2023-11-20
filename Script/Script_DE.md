# Verteilte Systeme

Die folgenden Texte sind Kapitel aus dem Buch "M. van Steen and A.S. Tanenbaum, Distributed Systems, 4th ed.,
distributed-systems.net, 2023".
Eine digitale Kopie kann gratis auf der Website des Buches heruntergeladen
werden: [distributed-systems.net](https://www.distributed-systems.net/index.php/books/ds4/).
Auf Deutsch ist nur die zweite Ausgabe verfügbar, welche aufgrund der Schnelllebigkeit der Materie nicht empfohlen wird.

Die folgenden Texte wurden mit Chat GPT-4 übersetzt und auf vom Dozenten überprüft.

# Einführung

Das Tempo, mit dem sich Computersysteme ändern, war, ist und bleibt überwältigend.
Von 1945, als die moderne Computerära begann, bis etwa 1985 waren Computer gross und teuer.
Da es keine Möglichkeit gab, sie zu verbinden, arbeiteten diese Computer unabhängig voneinander.

Ab Mitte der 1980er Jahre begannen jedoch zwei technologische Fortschritte, diese Situation zu ändern.
Das erste war die Entwicklung leistungsstarker Mikroprozessoren.
Anfangs handelte es sich um 8-Bit-Maschinen, aber bald wurden 16-, 32- und 64-Bit-CPUs üblich.
Mit leistungsstarken Multicore-CPUs stehen wir jetzt wieder vor der Herausforderung, Programme anzupassen und zu
entwickeln, um Parallelität zu nutzen.
Die aktuellen Maschinengenerationen haben die Rechenleistung der Mainframes von vor 30 oder 40 Jahren, kosten aber nur
einen Bruchteil davon.

Die zweite Entwicklung war die Erfindung von Hochgeschwindigkeits-Computernetzwerken.
Lokale Netzwerke oder LANs ermöglichen es, Tausende von Maschinen in einem Gebäude so zu verbinden, dass kleine
Informationsmengen in wenigen Mikrosekunden übertragen werden können.
Grössere Datenmengen können mit Raten von Milliarden Bits pro Sekunde (bps) zwischen Maschinen übertragen werden.
Wide Area Networks (WAN) ermöglichen die Verbindung von Hunderten von Millionen von Maschinen auf der ganzen Welt mit
Geschwindigkeiten von Zehntausenden bis zu Hunderten von Millionen bps und mehr.

Parallel zur Entwicklung immer leistungsfähigerer und vernetzter Maschinen konnten wir auch die Miniaturisierung von
Computersystemen beobachten, wobei das Smartphone vielleicht das beeindruckendste Ergebnis ist.
Mit Sensoren, viel Speicher und einer leistungsstarken Multicore-CPU sind diese Geräte nichts weniger als vollwertige
Computer.
Natürlich haben sie auch Netzwerkfähigkeiten. Ebenso sind sogenannte Nanocomputer leicht verfügbar geworden.
Diese kleinen Einplatinencomputer, oft nicht grösser als eine Kreditkarte, bieten problemlos nahezu Desktop-Leistung.
Bekannte Beispiele sind Raspberry Pi und Arduino Systeme.

Und die Geschichte geht weiter.
Da die Digitalisierung unserer Gesellschaft fortschreitet, werden wir uns immer bewusster, wie viele Computer
tatsächlich eingesetzt werden, regelmässig in andere Systeme eingebettet, wie Autos, Flugzeuge, Gebäude, Brücken, das
Stromnetz usw.
Dieses Bewusstsein wird leider verstärkt, wenn sich solche Systeme plötzlich als hackbar herausstellen. Zum Beispiel
wurde 2021 eine Treibstoffpipeline in den USA durch einen Ransomware-Angriff effektiv stillgelegt.
In diesem Fall bestand das Computersystem aus einer Mischung von Sensoren, Aktoren, Controllern, eingebetteten
Computern, Servern usw., die alle zu einem einzigen System zusammengeführt wurden.
Was viele von uns nicht realisieren, ist, dass lebenswichtige Infrastrukturen, wie Treibstoffpipelines, von vernetzten
Computersystemen überwacht und gesteuert werden.
In die gleiche Richtung gedacht, könnte es an der Zeit sein zu erkennen, dass ein modernes Auto eigentlich ein autonom
arbeitender, mobile vernetzter Computer ist.
In diesem Fall muss man sich nicht mit dem von einer Person getragenen mobilen Computer befassen, sondern mit dem
mobilen Computer, der Menschen befördert.

Die Grösse eines vernetzten Computersystems kann von einer Handvoll Geräten bis zu Millionen von Computern variieren.
Das Verbindungsnetzwerk kann verdrahtet, drahtlos oder eine Kombination aus beidem sein.
Darüber hinaus sind diese Systeme oft hochdynamisch, da Computer beitreten und verlassen können, wobei die Topologie und
Leistung des zugrunde liegenden Netzwerks fast ständig wechselt.

Es ist schwierig, sich Computersysteme vorzustellen, die nicht vernetzt sind.
Und tatsächlich können die meisten vernetzten Computersysteme von überall auf der Welt aus zugegriffen werden, weil sie
mit dem Internet verbunden sind.
Das Studium dieser Systeme kann leicht extrem komplex werden.
In diesem Kapitel beginnen wir, Licht darauf zu werfen, was verstanden werden muss, um das grössere Bild zu erstellen,
ohne sich zu verlieren.

## 1.1 Von vernetzten Systemen zu verteilten Systemen

Bevor wir uns verschiedenen Aspekten von verteilten Systemen zuwenden, sollten wir zunächst überlegen, was Verteilung
oder Dezentralisierung eigentlich bedeutet.

### 1.1.1 Verteilte versus dezentralisierte Systeme

Wenn man verschiedene Quellen betrachtet, gibt es viele Meinungen über verteilte versus dezentralisierte Systeme.
Oft wird der Unterschied durch drei verschiedene Organisationen von vernetzten Computersystemen veranschaulicht.
Hierbei stellt jeder Knoten ein Computersystem dar und eine Kante einen Kommunikationslink zwischen zwei Knoten.

Inwieweit solche Unterscheidungen nützlich sind, muss noch geprüft werden, insbesondere wenn Diskussionen über die Vor-
und Nachteile jeder Organisation beginnen. Zum Beispiel wird oft behauptet, dass zentralisierte Organisationen nicht gut
skalieren. Ebenso wird gesagt, dass verteilte Organisationen robuster gegenüber Ausfällen sind. Wie wir sehen werden,
sind diese Behauptungen generell nicht wahr.

Wir gehen einen anderen Weg. Wenn wir von einem vernetzten Computersystem als einer Sammlung von Computern in einem
Netzwerk ausgehen, können wir uns fragen, wie diese Computer überhaupt miteinander verbunden wurden. Es gibt zwei
Sichtweisen, die man einnehmen kann.

Die erste, **integrative Sichtweise**, ist, dass es notwendig war, bestehende (vernetzte) Computersysteme miteinander zu
verbinden. Dies geschieht in der Regel, wenn Dienste, die auf einem System laufen, für Benutzer und Anwendungen
verfügbar gemacht werden müssen, an die zuvor nicht gedacht wurde. Ein Beispiel dafür ist die Integration von
Finanzdienstleistungen mit Projektmanagementdiensten, wie es oft in einer einzelnen Organisation der Fall ist. Im
wissenschaftlichen Forschungsbereich haben wir Bemühungen gesehen, eine Vielzahl von oft teuren Ressourcen (
Spezialcomputer, Supercomputer, sehr grosse Datenbanksysteme usw.) zu dem zu verbinden, was als Grid-Computer bezeichnet
wird.

Die zweite, **expansive Sichtweise** ist, dass ein bestehendes System durch zusätzliche Computer erweitert werden
musste. Diese Sichtweise steht im Zusammenhang mit dem Gebiet der verteilten Systeme. Es geht darum, ein System mit
Computern zu erweitern, um Ressourcen dort bereitzustellen, wo sie benötigt werden. Eine Erweiterung kann auch durch das
Bedürfnis nach grösserer Zuverlässigkeit getrieben werden: Wenn ein Computer ausfällt, gibt es andere, die übernehmen
können. Eine wichtige Art von Erweiterung ist, wenn ein Dienst für entfernte Benutzer und Anwendungen verfügbar gemacht
werden muss, zum Beispiel durch Bereitstellung einer Web-Schnittstelle oder einer Smartphone-Anwendung. Dieses letzte
Beispiel zeigt auch, dass die Unterscheidung zwischen einer integrativen und einer expansiven Sichtweise nicht eindeutig
ist.

In beiden Fällen sehen wir, dass das vernetzte System Dienste ausführt, wobei jeder Dienst als eine Sammlung von
Prozessen und Ressourcen implementiert ist, die sich über mehrere Computer erstrecken. Die beiden Ansichten führen zu
einer natürlichen Unterscheidung zwischen zwei Arten von vernetzten Computersystemen:

- Ein dezentralisiertes System ist ein vernetztes Computersystem, bei dem Prozesse und Ressourcen **notwendigerweise**
  auf mehrere Computer verteilt sind.
- Ein verteiltes System ist ein vernetztes Computersystem, bei dem Prozesse und Ressourcen **ausreichend** auf mehrere
  Computer verteilt sind.

Bevor wir diskutieren, warum diese Unterscheidung wichtig ist, betrachten wir einige Beispiele für jeden Systemtyp.

Dezentralisierte Systeme beziehen sich hauptsächlich auf die integrative Sichtweise von vernetzten Computersystemen.
Sie entstehen, weil wir Systeme verbinden wollen, aber möglicherweise durch administrative Grenzen eingeschränkt werden.
Viele Anwendungen im Bereich der künstlichen Intelligenz benötigen beispielsweise riesige Datenmengen, um zuverlässige
Vorhersagemodelle zu erstellen.
Normalerweise werden Daten zu den Hochleistungscomputern gebracht, die buchstäblich Modelle trainieren, bevor sie
verwendet werden können.
Wenn Daten jedoch innerhalb des Perimeters einer Organisation bleiben müssen (und es gibt viele Gründe, warum dies
notwendig ist), müssen wir das Training zu den Daten bringen.
Das Ergebnis ist als föderiertes Lernen bekannt und wird durch ein dezentralisiertes System implementiert, bei dem das
Bedürfnis nach Verteilung von Prozessen und Ressourcen durch administrative Richtlinien vorgegeben wird.

Ein weiteres Beispiel für ein dezentralisiertes System ist das eines Distributed-Ledgers (verteiltes Kassenbuch), auch
bekannt als
Blockchain.
In diesem Fall müssen wir mit der Situation umgehen, dass sich die teilnehmenden Parteien nicht genug vertrauen, um
einfache Kooperationsschemata aufzustellen.
Stattdessen machen sie im Wesentlichen Transaktionen untereinander vollständig öffentlich (und verifizierbar) durch
einen
nur erweiterbaren Ledger, der Aufzeichnungen dieser Transaktionen enthält.
Der Ledger selbst ist vollständig über die Teilnehmer verteilt, und die Teilnehmer sind es, die Transaktionen (anderer)
validieren, bevor sie sie dem Ledger hinzufügen.
Das Ergebnis ist ein dezentralisiertes System, bei dem Prozesse und Ressourcen tatsächlich notwendigerweise auf mehrere
Computer verteilt sind, in diesem Fall aufgrund fehlenden Vertrauens.
Als letztes Beispiel für ein dezentralisiertes System betrachten Sie Systeme, die von Natur aus geografisch verteilt
sind.
Dies tritt typischerweise bei Systemen auf, bei denen ein tatsächlicher Ort überwacht werden muss, beispielsweise im
Falle eines Kraftwerks, eines Gebäudes, einer bestimmten natürlichen Umgebung usw.
Das System, das die Überwachung steuert und Entscheidungen trifft, kann leicht woanders platziert werden als der
überwachte Ort.
Ein offensichtliches Beispiel ist die Überwachung und Steuerung von Satelliten, aber auch alltäglichere Situationen wie
die Überwachung und Steuerung des Verkehrs, von Zügen usw.
In diesen Beispielen ergibt sich die Notwendigkeit zur Verteilung von Prozessen und Ressourcen aus einem räumlichen
Argument.

Wie wir bereits erwähnt haben, stehen verteilte Systeme hauptsächlich im Zusammenhang mit der weitreichenden Sicht auf
vernetzte Computersysteme.
Ein bekanntes Beispiel ist die Nutzung von E-Mail-Diensten, wie Google Mail.
Was oft passiert, ist, dass ein Benutzer sich über eine Web-Schnittstelle in das System einloggt, um Mails zu lesen und
zu senden.
Häufiger ist jedoch, dass Benutzer ihren persönlichen Computer (wie zum Beispiel einen Laptop) konfigurieren, um einen
bestimmten Mail-Client zu nutzen.
Dafür müssen sie einige Einstellungen vornehmen, wie den eingehenden und ausgehenden Server.
Im Falle von Google Mail sind dies imap.gmail.com und smtp.gmail.com.
Logischerweise scheint es, als würden diese beiden Server alle Ihre Mails bearbeiten.
Mit einer Schätzung von fast 2 Milliarden Benutzern im Jahr 2022 ist es jedoch unwahrscheinlich, dass nur zwei Computer
all ihre E-Mails bearbeiten können (die auf mehr als 300 Milliarden pro Jahr geschätzt wurden, das heisst, etwa 10.000
Mails pro Sekunde).
Hinter den Kulissen wurde der gesamte Google Mail-Dienst natürlich auf viele Computer verteilt, die gemeinsam ein
verteiltes System bilden.
Dieses System wurde eingerichtet, um sicherzustellen, dass so viele Benutzer ihre Mails verarbeiten können (d.h. es
gewährleistet die Skalierbarkeit), aber auch, dass das Risiko des Mailverlusts aufgrund von Ausfällen minimal ist (d.h.
das System gewährleistet Fehlertoleranz).
Für den Benutzer wird jedoch das Bild von nur zwei Servern aufrechterhalten (d.h. die Verteilung selbst ist für den
Benutzer hochgradig transparent).
Das verteilte System, das einen E-Mail-Dienst wie Google Mail implementiert, erweitert (oder verkleinert) sich je nach
den Anforderungen an die Zuverlässigkeit, die wiederum von der Anzahl seiner Benutzer abhängt.

Ein völlig anderer Typ von verteiltem System wird durch die Sammlung sogenannter Content Delivery Networks oder kurz
CDNs gebildet.
Ein bekanntes Beispiel ist Akamai mit, im Jahr 2022, über 400.000 Servern weltweit.
Das grundsätzliche Funktionieren von CDNs werden wir später in Kapitel 3 besprechen.
Im Grunde genommen wird der Inhalt einer tatsächlichen Website kopiert und auf verschiedene Server des CDN verteilt.
Wenn ein Benutzer eine Website besucht, wird er transparent zu einem nahegelegenen Server umgeleitet, der den gesamten
oder einen Teil des Inhalts dieser Website enthält.
Die Auswahl des Servers kann von vielen Dingen abhängen, sicherlich aber bei Streaming-Inhalten wird ein Server
ausgewählt, für den eine gute Leistung in Bezug auf Latenz und Bandbreite gewährleistet werden kann.
Das CDN stellt dynamisch sicher, dass der ausgewählte Server den erforderlichen Inhalt sofort verfügbar hat,
aktualisiert diesen Inhalt bei Bedarf oder entfernt ihn vom Server, wenn dort keine oder nur sehr wenige Benutzer
bedient werden müssen.
Der Benutzer weiss derweil nichts darüber, was hinter den Kulissen vor sich geht (was wiederum eine Form der
Verteilungstransparenz ist).
In diesem Beispiel sehen wir auch, dass Inhalte nicht auf alle Server kopiert werden, sondern nur dorthin, wo es
sinnvoll ist, also ausreichend und aus Leistungsgründen.
CDNs kopieren Inhalte auch auf mehrere Server, um ein hohes Mass an Zuverlässigkeit zu gewährleisten.

Als abschliessendes, viel kleineres verteiltes System betrachten Sie eine Einrichtung basierend auf einem
Network-Attached Storage-Gerät, auch NAS genannt.
Für den Hausgebrauch besteht ein typisches NAS aus 2–4 Steckplätzen für interne Festplatten.
Das NAS fungiert als Dateiserver: Es ist über ein (in der Regel drahtloses) Netzwerk für jedes autorisierte Gerät
zugänglich und kann so Dienste wie gemeinsam genutzten Speicher, automatisierte Backups, Streaming-Medien usw. anbieten.
Das NAS selbst kann am besten als ein einzelner, für die Speicherung von Dateien optimierter Computer gesehen werden,
der die Möglichkeit bietet, diese Dateien leicht zu teilen.
Letzteres ist wichtig, und zusammen mit mehreren Benutzern haben wir im Wesentlichen eine Einrichtung eines verteilten
Systems.
Die Benutzer arbeiten mit einer Dateisammlung, die lokal (d.h. von ihrem Laptop aus) leicht zugänglich ist (tatsächlich
scheinbar in das lokale Dateisystem integriert), während sie auch direkt von anderen Benutzern zugänglich ist.
Wiederum ist verborgen, wo und wie die geteilten Dateien gespeichert sind (d.h. die Verteilung ist transparent).
Wenn man annimmt, dass das Teilen von Dateien das Ziel ist, dann sehen wir, dass ein NAS tatsächlich eine ausreichende
Verteilung von Prozessen und Ressourcen bieten kann.

### 1.1.2 Warum die Unterscheidung relevant ist

Warum machen wir diese Unterscheidung zwischen dezentralen und verteilten Systemen?
Es ist wichtig zu erkennen, dass zentralisierte Lösungen generell viel einfacher sind, und das auch entlang
unterschiedlicher Kriterien.
Dezentralisierung, also das Verteilen der Implementierung eines Dienstes über mehrere Computer, weil wir glauben, dass
es notwendig ist, ist eine Entscheidung, die sorgfältig überdacht werden muss.
Tatsächlich sind verteilte und dezentralisierte Lösungen von Natur aus schwierig:

- Es gibt viele, oft unerwartete, Abhängigkeiten, die das Verständnis des Verhaltens dieser Systeme behindern.
- Verteilte und dezentralisierte Systeme leiden fast ständig unter partiellen Ausfällen: Ein Prozess oder eine Ressource
  an einem der beteiligten Computer funktioniert nicht wie erwartet. Es kann tatsächlich eine Weile dauern, diesen
  Fehler zu entdecken, während solche Fehler vorzugsweise maskiert werden (d.h., sie bleiben für Benutzer und
  Anwendungen unbemerkt), einschliesslich der Wiederherstellung von Ausfällen.
- In engem Zusammenhang mit partiellen Ausfällen steht die Tatsache, dass in vielen vernetzten Computersystemen
  teilnehmende Knoten, Prozesse, Ressourcen usw. kommen und gehen. Dies macht diese Systeme hochdynamisch und erfordert
  Formen des automatisierten Managements und der Wartung, was wiederum die Komplexität erhöht.
- Die Tatsache, dass verteilte und dezentralisierte Systeme vernetzt sind, von vielen Benutzern und Anwendungen genutzt
  werden und oft mehrere administrative Grenzen überschreiten, macht sie besonders anfällig für Sicherheitsangriffe.
  Daher erfordert das Verständnis dieser Systeme und ihres Verhaltens, dass wir verstehen, wie sie gesichert werden
  können und wie sie tatsächlich gesichert sind. Leider ist das Verständnis von Sicherheit nicht so einfach.

Unsere Unterscheidung liegt zwischen der Ausreichendkeit und der Notwendigkeit, Prozesse und Ressourcen über mehrere
Computer zu verteilen.
In diesem Buch gehen wir davon aus, dass Dezentralisierung niemals ein Selbstzweck sein kann und dass sie sich auf die
Ausreichendkeit zur Verteilung von Prozessen und Ressourcen über Computer konzentrieren sollte.
Grundsätzlich gilt: Je weniger Verteilung, desto besser.
Gleichzeitig müssen wir jedoch erkennen, dass Verteilung manchmal wirklich notwendig ist, wie die Beispiele für
dezentralisierte Systeme zeigen.
Aus diesem Blickwinkel der Ausreichendkeit geht es in diesem Buch primär um verteilte Systeme, und wo es angebracht ist,
werden dezentralisierte Systemen besprochen.

In dieselbe Richtung gehend, in Anbetracht der Tatsache, dass verteilte und dezentralisierte Systeme von Natur aus
komplex sind, ist es ebenso wichtig, Lösungen in Betracht zu ziehen, die so einfach wie möglich sind.
Daher werden wir Optimierungen von Lösungen kaum diskutieren, fest davon überzeugt, dass der Leistungsgewinn selten die
erhöhte Komplexität rechtfertigt.

### 1.1.3 Das Studium verteilter Systeme

In Anbetracht der Tatsache, dass verteilte Systeme von Natur aus schwierig sind, ist es wichtig, einen systematischen
Ansatz zu ihrer Untersuchung zu wählen.
Eine unserer grössten Sorgen ist, dass es in verteilten Systemen so viele explizite und implizite Abhängigkeiten gibt.

Zum Beispiel gibt es so etwas wie ein separates Kommunikationsmodul oder ein separates Sicherheitsmodul nicht.
Unser Ansatz besteht darin, verteilte Systeme aus einer begrenzten Anzahl, aber verschiedenen Perspektiven zu
betrachten.
Jede Perspektive wird in einem separaten Kapitel betrachtet.

- Es gibt viele Möglichkeiten, wie verteilte Systeme organisiert sind. Wir beginnen unsere Diskussion mit der
  architektonischen Perspektive: Was sind gängige Organisationen, was sind gängige Stile? Die architektonische
  Perspektive wird helfen, einen ersten Eindruck davon zu bekommen, wie verschiedene Komponenten bestehender Systeme
  interagieren und voneinander abhängen.
- Bei verteilten Systemes dreht sich alles um Prozesse. Die Prozessperspektive befasst sich voll und ganz mit dem
  Verständnis der verschiedenen Formen von Prozessen, die in verteilten Systemen auftreten, sei es Threads, ihre
  Virtualisierung von Hardware-Prozessen, Clients, Server und so weiter. Prozesse bilden das Software-Rückgrat
  verteilter Systeme, und ihr Verständnis ist für das Verständnis verteilter Systeme unerlässlich.
- Offensichtlich ist bei mehreren Computern die Kommunikation zwischen Prozessen unerlässlich. Die
  Kommunikationsperspektive betrifft die Einrichtungen, die verteilte Systeme bereitstellen, um Daten zwischen Prozessen
  auszutauschen. Sie beinhaltet im Wesentlichen Prozeduraufrufen über mehrere Computer, Messaging mit einer Vielzahl von
  semantischen Optionen und verschiedene Arten der Kommunikation zwischen Prozessgruppen.
- Damit verteilte Systeme funktionieren, koordinieren die Prozesse Dinge unter der Haube, auf der Anwendungen ausgeführt
  werden. Sie koordinieren gemeinsam, zum Beispiel, um den Mangel an globaler Uhr auszugleichen, um exklusiven Zugriff
  auf gemeinsame Ressourcen zu realisieren und so weiter. Die Koordinationsperspektive beschreibt eine Reihe
  grundlegender Koordinationsaufgaben, die als Teil der meisten verteilten Systeme durchgeführt werden müssen.
- Um auf Prozesse und Ressourcen zuzugreifen, benötigen wir Namensgebung. Insbesondere benötigen wir Namensschemas, die
  bei Verwendung zum Prozess, zu Ressourcen oder zu welcher anderen Art von Entität auch immer führen, die benannt wird.
  So einfach das auch erscheinen mag, stellt sich die Namensgebung in verteilten Systemen nicht nur als entscheidend
  heraus, es gibt auch viele Möglichkeiten, wie die Namensgebung unterstützt wird. Die Namensperspektive konzentriert
  sich voll und ganz darauf, einen Namen zum Zugriffspunkt der benannten Entität aufzulösen.
- Ein kritischer Aspekt verteilter Systeme ist, dass sie sowohl in Bezug auf die Effizienz als auch in Bezug auf die
  Zuverlässigkeit gut funktionieren. Das Hauptinstrument für beide Aspekte ist die Replikation von Ressourcen. Das
  einzige Problem mit der Replikation ist, dass Updates stattfinden können, was bedeutet, dass alle Kopien einer
  Ressource ebenfalls aktualisiert werden müssen. Hier wird es herausfordernd, den Anschein eines nicht verteilten
  Systems aufrechtzuerhalten. Die Konsistenz- und Replikationsperspektive konzentriert sich im Wesentlichen auf die
  Kompromisse zwischen Konsistenz, Replikation und Leistung.
- Wie bereits erwähnt, sind verteilte Systeme partiellen Ausfällen unterworfen. Die Perspektive der Fehlertoleranz
  taucht in die Mittel zur Maskierung von Ausfällen und ihrer Wiederherstellung ein. Es hat sich als eine der
  schwierigsten Perspektiven für das Verständnis verteilter Systeme erwiesen, hauptsächlich weil es so viele Kompromisse
  zu machen gibt und weil es nachweislich unmöglich ist, Ausfälle und deren Wiederherstellung vollständig zu maskieren.
- Wie ebenfalls erwähnt, gibt es so etwas wie ein nicht gesichertes verteiltes System nicht. Die Sicherheitsperspektive
  konzentriert sich darauf, wie man den autorisierten Zugriff auf Ressourcen sicherstellt. Zu diesem Zweck müssen wir
  das Vertrauen in verteilten Systemen sowie die Authentifizierung, nämlich die Überprüfung einer behaupteten Identität,
  diskutieren. Die Sicherheitsperspektive kommt zuletzt, doch später in diesem Kapitel werden wir einige grundlegende
  Instrumente diskutieren, die benötigt werden, um die Rolle der Sicherheit in den vorherigen Perspektiven zu verstehen.

## 1.2 Entwurfsziele

Nur weil es möglich ist, verteilte Systeme zu bauen, bedeutet das nicht zwangsläufig, dass es eine gute Idee ist.
In diesem Abschnitt diskutieren wir vier wichtige Ziele, die erfüllt werden sollten, um den Aufbau eines verteilten
Systems
zu rechtfertigen.
Ein verteiltes System sollte Ressourcen leicht zugänglich machen;
es sollte die Tatsache verbergen, dass Ressourcen über ein Netzwerk verteilt sind;
es sollte offen sein;
und es sollte skalierbar sein.

### 1.2.1 Ressourcenfreigabe

Ein wichtiges Ziel eines verteilten Systems ist es, Benutzern (und Anwendungen) den einfachen Zugriff auf und die
gemeinsame Nutzung von entfernten Ressourcen zu ermöglichen. Ressourcen können praktisch alles sein, typische Beispiele
sind jedoch Peripheriegeräte, Speichereinrichtungen, Daten, Dateien, Dienste und Netzwerke, um nur einige zu nennen. Es
gibt viele Gründe, Ressourcen teilen zu wollen. Ein offensichtlicher Grund ist der der Wirtschaftlichkeit. Es ist
beispielsweise billiger, eine einzelne hochwertige zuverlässige Speichereinrichtung gemeinsam zu nutzen, als für jeden
Benutzer separat Speicher zu kaufen und zu warten. Das Verbinden von Benutzern und Ressourcen erleichtert auch die
Zusammenarbeit und den Informationsaustausch, wie der Erfolg des Internets mit seinen einfachen Protokollen zum
Austausch von Dateien, Mail, Dokumenten, Audio und Video zeigt. Die Konnektivität des Internets hat es geografisch weit
verstreuten Personengruppen ermöglicht, durch alle Arten von Groupware zusammenzuarbeiten, das heisst Software für
kollaboratives Bearbeiten, Videokonferenzen und so weiter, wie es beispielsweise von multinationalen
Softwareentwicklungsunternehmen illustriert wird, die einen Grossteil ihrer Codeproduktion nach Asien ausgelagert haben,
aber auch durch die Vielzahl von Zusammenarbeitstools, die durch die COVID-19-Pandemie (leichter) verfügbar wurden.

Die Ressourcenfreigabe in verteilten Systemen wird auch durch den Erfolg von File-Sharing-Peer-to-Peer-Netzwerken wie
BitTorrent illustriert. Diese verteilten Systeme ermöglichen es Benutzern, Dateien einfach über das Internet zu teilen.
Peer-to-Peer-Netzwerke werden oft mit der Verteilung von Medien Dateien wie Audio und Video in Verbindung gebracht. In
anderen Fällen wird die Technologie zum Verteilen grosser Datenmengen verwendet, wie im Fall von Software-Updates,
Backup-Diensten und Datenabgleich über mehrere Server.

Die nahtlose Integration von Ressourcenfreigabeeinrichtungen in einer vernetzten Umgebung ist ebenfalls heute üblich.
Eine Benutzergruppe kann einfach Dateien in einen speziellen gemeinsamen Ordner legen, der von einem Drittanbieter
irgendwo im Internet gepflegt wird.
Mit spezieller Software ist der gemeinsam genutzte Ordner kaum von anderen Ordnern auf dem Computer eines Benutzers zu
unterscheiden.
Tatsächlich ersetzen diese Dienste die Verwendung eines gemeinsamen Verzeichnisses in einem lokalen verteilten
Dateisystem und
machen Daten unabhängig von der Organisation, der sie angehören, und unabhängig von ihrem Aufenthaltsort verfügbar. Der
Dienst wird für verschiedene Betriebssysteme angeboten. Wo genau Daten gespeichert werden, ist für den Endbenutzer
völlig verborgen.

### 1.2.2 Verteilungstransparenz

Ein wichtiges Ziel eines verteilten Systems ist es, die Tatsache zu verbergen, dass seine Prozesse und Ressourcen
physisch über mehrere Computer verteilt sind, möglicherweise über grosse Entfernungen hinweg. Mit anderen Worten, es
versucht, die Verteilung von Prozessen und Ressourcen für Endbenutzer und Anwendungen transparent, das heisst unsichtbar,
zu machen. Wie wir in Kapitel 2 ausführlicher diskutieren werden, wird die Verteilungstransparenz durch die sogenannte
Middleware realisiert, skizziert in Abbildung 1.2 (siehe Gazis und Katsiri [2022] für eine erste Einführung). Im
Wesentlichen sehen Anwendungen überall die gleiche Schnittstelle, während hinter dieser Schnittstelle, wo und wie
Prozesse und Ressourcen sind und wie sie zugegriffen werden, transparent gehalten wird. Es gibt verschiedene Arten von
Transparenz, die wir als Nächstes besprechen.

#### Arten von Verteilungstransparenz

Das Konzept der Transparenz kann auf mehrere Aspekte eines verteilten Systems angewendet werden, von denen die
wichtigsten in Abbildung 1.3 aufgeführt sind. Wir verwenden den Begriff Objekt, um entweder einen Prozess oder eine
Ressource zu bezeichnen.

**Zugriffstransparenz** befasst sich damit, Unterschiede in der Datenrepräsentation und der Art und Weise, wie Objekte
zugegriffen werden kann, zu verbergen. Auf grundlegender Ebene möchten wir Unterschiede in Maschinenarchitekturen
verbergen, aber wichtiger ist, dass wir uns darauf einigen, wie Daten von verschiedenen Maschinen und Betriebssystemen
dargestellt werden sollen. Ein verteiltes System könnte beispielsweise Computersysteme haben, die verschiedene
Betriebssysteme ausführen, von denen jedes seine eigenen Dateibenennungskonventionen hat. Unterschiede in den
Benennungskonventionen, Unterschiede in Dateioperationen oder Unterschiede in der Art und Weise, wie die Kommunikation
auf niedriger Ebene mit anderen Prozessen stattfinden soll, sind Beispiele für Zugriffsprobleme, die vorzugsweise vor
Benutzern und Anwendungen verborgen bleiben sollten.

**Ortstransparenz** betrifft die Tatsache, dass die physische Position eines Objekts für Benutzer und Anwendungen nicht
offensichtlich sein sollte. In einem verteilten System ist es beispielsweise unerheblich, ob eine Datei auf dem lokalen
Rechner oder auf einem Rechner am anderen Ende der Welt gespeichert ist. In beiden Fällen sollte der Zugriff auf die
Datei auf die gleiche Weise erfolgen, ohne dass der Benutzer sich der tatsächlichen Position bewusst ist.

**Migrationstransparenz** bezieht sich darauf, dass die Verlagerung (Migration) von Ressourcen und Prozessen innerhalb
eines Systems vor Benutzern und Anwendungen verborgen bleiben sollte. Dies ist besonders wichtig in Systemen, die
Ressourcen je nach Anforderung und Verfügbarkeit dynamisch verlagern können.

**Replikationstransparenz** bedeutet, dass die Duplikation von Prozessen oder Ressourcen in einem System, um etwa die
Zuverlässigkeit oder Performance zu erhöhen, für Benutzer und Anwendungen unsichtbar sein sollte.

**Parallelitätstransparenz** besagt, dass die gleichzeitige Ausführung von Prozessen oder die gleichzeitige Nutzung von
Ressourcen für Benutzer und Anwendungen unsichtbar sein sollte.

Zuletzt, aber sicherlich nicht am wenigsten wichtig, ist es von Bedeutung, dass ein verteiltes System eine
**Fehlertransparenz** bietet. Dies bedeutet, dass ein Benutzer oder eine Anwendung nicht bemerkt, wenn ein Teil des
Systems nicht ordnungsgemäss funktioniert, und dass das System anschliessend (und automatisch) von diesem Fehler erholt.
Fehler zu maskieren ist eines der schwierigsten Probleme in verteilten Systemen und ist sogar unmöglich, wenn bestimmte
scheinbar realistische Annahmen getroffen werden, wie wir in Kapitel 8 besprechen werden. Die Hauptproblematik beim
Maskieren und transparenten Wiederherstellen von Fehlern liegt in der Unfähigkeit, zwischen einem "toten" Prozess und
einem schmerzhaft langsam reagierenden zu unterscheiden. Zum Beispiel wird ein Browser beim Kontaktieren eines
beschäftigten Web-Servers schliesslich eine Zeitüberschreitung melden und berichten, dass die Webseite nicht verfügbar
ist. Zu diesem Zeitpunkt kann der Benutzer nicht feststellen, ob der Server tatsächlich ausgefallen ist oder das
Netzwerk stark überlastet ist.

#### Grad der Verteilungstransparenz

Obwohl Verteilungstransparenz allgemein für jedes verteilte System als bevorzugt angesehen wird, gibt es Situationen, in
denen es keine gute Idee ist, blindlings zu versuchen, alle Verteilungsaspekte vor den Benutzern zu verbergen. Ein
einfaches Beispiel ist die Anforderung, dass Ihre elektronische Zeitung wie gewohnt vor 7 Uhr morgens in Ihrem Postfach
erscheint, während Sie sich gerade am anderen Ende der Welt in einer anderen Zeitzone befinden. Ihre Morgenzeitung wird
nicht die gewohnte Morgenzeitung sein.

Ebenso kann von einem weitverteilten System, das einen Prozess in San Francisco mit einem Prozess in Amsterdam
verbindet, nicht erwartet werden, dass es die Tatsache verbirgt, dass Mutter Natur es nicht zulässt, eine Nachricht von
einem Prozess zum anderen in weniger als ungefähr 35 Millisekunden zu senden. Die Praxis zeigt, dass dies mit einem
Computernetzwerk tatsächlich mehrere hundert Millisekunden dauert. Die Signalübertragung wird nicht nur durch die
Lichtgeschwindigkeit begrenzt, sondern auch durch begrenzte Verarbeitungskapazitäten und Verzögerungen in den
Zwischenstationen.

Es gibt auch einen Trade-off zwischen einem hohen Grad an Transparenz und der Leistung eines Systems. Zum Beispiel
versuchen viele Internetanwendungen wiederholt, einen Server zu kontaktieren, bevor sie schliesslich aufgeben. Das
Versuchen, einen vorübergehenden Serverausfall zu maskieren, bevor ein anderer versucht wird, kann das gesamte System
verlangsamen. In einem solchen Fall wäre es vielleicht besser gewesen, früher aufzugeben oder zumindest dem Benutzer zu
erlauben, die Kontaktversuche abzubrechen.

Ein weiteres Beispiel ist, wo wir garantieren müssen, dass mehrere Replikate, die auf verschiedenen Kontinenten
befindlich sind, jederzeit konsistent sein müssen. Das heisst, wenn eine Kopie geändert wird, sollte diese Änderung zu
allen Kopien propagiert werden, bevor eine andere Operation erlaubt wird. Ein einziges Update kann jetzt sogar Sekunden
dauern, etwas, das den Benutzern nicht verborgen werden kann.

Schliesslich gibt es Situationen, in denen es überhaupt nicht offensichtlich ist, dass das Verbergen der Verteilung eine
gute Idee ist.
Da sich verteilte Systeme auf Geräte ausweiten, die Menschen herumtragen, und das Standort- und Kontextbewusstsein immer
wichtiger wird, ist es vielleicht am besten, die Verteilung offenzulegen, anstatt zu versuchen, sie zu verbergen.
Ein offensichtliches Beispiel ist die Nutzung von standortbasierten Diensten, die oft auf Mobiltelefonen zu finden sind,
wie z. B. das Finden eines nächstgelegenen Geschäfts oder nahen Freunden.
Hier ist es explizit erwünscht, dass die Standorte ersichtlich sind und nicht verborgen werden.

Es gibt auch andere Argumente gegen die Verteilungstransparenz. Da wir erkennen, dass vollständige
Verteilungstransparenz einfach unmöglich ist, sollten wir uns fragen, ob es überhaupt klug ist, so zu tun, als könnten
wir sie erreichen. Es könnte viel besser sein, die Verteilung explizit zu machen, damit der Benutzer und der
Anwendungsentwickler nie in die Irre geführt werden, zu glauben, dass es so etwas wie Transparenz gibt. Das Ergebnis
wird sein, dass die Benutzer das (manchmal unerwartete) Verhalten eines verteilten Systems viel besser verstehen und
somit viel besser darauf vorbereitet sind, mit diesem Verhalten umzugehen.

Der Schluss ist, dass das Anstreben von Verteilungstransparenz beim Entwerfen und Implementieren von verteilten Systemen
ein gutes Ziel sein kann, aber dass es zusammen mit anderen Fragen wie Leistung und Verständlichkeit betrachtet werden
sollte. Der Aufwand für das Erreichen vollständiger Transparenz könnte überraschend hoch sein.

### 1.2.3 Offenheit

Ein weiteres wichtiges Ziel verteilter Systeme ist die Offenheit. Ein offenes verteiltes System ist im Grunde ein
System, das Komponenten anbietet, die leicht von anderen Systemen verwendet oder in diese integriert werden können.
Gleichzeitig besteht ein offenes verteiltes System oft aus Komponenten, die von anderswo stammen.

#### Interoperabilität, Zusammensetzbarkeit und Erweiterbarkeit

Offen zu sein bedeutet, dass die Komponenten standardisierten Regeln folgen sollten, die die Syntax und Semantik
beschreiben, die diese Komponenten bieten (d.h., welchen Dienst sie erbringen). Ein allgemeiner Ansatz besteht darin,
Dienste über Schnittstellen mit einer Interface Definition Language (IDL) zu definieren. In IDL geschriebene
Schnittstellendefinitionen erfassen fast immer nur die Syntax der Dienste. Mit anderen Worten, sie spezifizieren genau
die Namen der verfügbaren Funktionen zusammen mit den Typen der Parameter, Rückgabewerte, möglichen Ausnahmen usw. Der
schwierige Teil ist die genaue Spezifikation dessen, was diese Dienste tun, d.h. die Semantik der Schnittstellen. In der
Praxis werden solche Spezifikationen oft informell in natürlicher Sprache gegeben.

Wenn richtig spezifiziert, ermöglicht eine Schnittstellendefinition einem beliebigen Prozess, der eine bestimmte
Schnittstelle benötigt, mit einem anderen Prozess zu kommunizieren, der diese Schnittstelle bereitstellt. Es erlaubt
auch zwei unabhängigen Parteien, völlig unterschiedliche Implementierungen dieser Schnittstellen zu erstellen, was zu
zwei separaten Komponenten führt, die genau gleich funktionieren.

Korrekte Spezifikationen sind vollständig und neutral. Vollständig bedeutet, dass alles, was zur Implementierung
notwendig ist, tatsächlich spezifiziert wurde. Viele Schnittstellendefinitionen sind jedoch überhaupt nicht vollständig,
sodass es notwendig ist, implementierungsspezifische Details hinzuzufügen. Ebenso wichtig ist, dass Spezifikationen
nicht vorschreiben, wie eine Implementierung aussehen sollte; sie sollten neutral sein.

Wie Blair und Stefani [1998] feststellten, sind Vollständigkeit und Neutralität wichtig für Interoperabilität und
Portabilität. Interoperabilität charakterisiert das Mass, in dem zwei Implementierungen von Systemen oder Komponenten
unterschiedlicher Hersteller koexistieren und lediglich auf die Dienste des anderen vertrauen können, wie sie durch
einen gemeinsamen Standard spezifiziert werden. Portabilität beschreibt, inwieweit eine für ein verteiltes System A
entwickelte Anwendung ohne Änderung auf einem anderen verteilten System B ausgeführt werden kann, das die gleichen
Schnittstellen wie A implementiert.

Ein weiteres wichtiges Ziel für ein offenes verteiltes System ist, dass es einfach sein sollte, das System aus
verschiedenen Komponenten (möglicherweise von verschiedenen Entwicklern) zu konfigurieren. Es sollte auch einfach sein,
neue Komponenten hinzuzufügen oder vorhandene zu ersetzen, ohne die verbleibenden Komponenten zu beeinflussen. Mit
anderen Worten, ein offenes verteiltes System sollte auch erweiterbar sein. Zum Beispiel sollte es in einem
erweiterbaren System relativ einfach sein, Teile hinzuzufügen, die auf einem anderen Betriebssystem laufen, oder sogar
ein gesamtes Dateisystem zu ersetzen. Relativ einfache Beispiele für Erweiterbarkeit sind Plug-ins für Webbrowser, aber
auch solche für Websites, wie sie für WordPress verwendet werden.

#### Trennung von Richtlinie und Mechanismus

Um Flexibilität in offenen verteilten Systemen zu erreichen, ist es entscheidend, dass das System als Sammlung von
relativ kleinen und leicht austauschbaren oder anpassbaren Komponenten organisiert wird. Dies impliziert, dass wir nicht
nur Definitionen für die höchsten Schnittstellen bereitstellen sollten, sondern auch Definitionen für Schnittstellen zu
internen Teilen des Systems und beschreiben, wie diese Teile interagieren. Dieser Ansatz ist relativ neu. Viele ältere
und auch zeitgenössische Systeme werden mit einem monolithischen Ansatz konstruiert, bei dem Komponenten nur logisch
getrennt, aber als ein riesiges Programm implementiert sind. Dieser Ansatz erschwert den Austausch oder die Anpassung
einer Komponente, ohne das gesamte System zu beeinflussen. Monolithische Systeme neigen daher dazu, geschlossen statt
offen zu sein.

Der Bedarf, ein verteiltes System zu ändern, wird oft durch eine Komponente verursacht, die nicht die optimale
Richtlinie für einen bestimmten Benutzer oder eine bestimmte Anwendung bereitstellt. Als Beispiel betrachten Sie das
Caching in Webbrowsern. Es müssen viele verschiedene Parameter berücksichtigt werden:
**Speicher**: Wo sollen Daten zwischengespeichert werden? Typischerweise gibt es neben dem Speicher auf der Festplatte
einen Cache im Speicher. In diesem Fall muss die genaue Position im lokalen Dateisystem berücksichtigt werden.
Ausnahme: Wenn der Cache voll ist, welche Daten sollen entfernt werden, damit neu abgerufene Seiten gespeichert werden
können?
**Teilen**: Verwendet jeder Browser einen privaten Cache, oder soll ein Cache zwischen Browsern verschiedener Benutzer
geteilt werden?
**Aktualisierung**: Wann überprüft ein Browser, ob zwischengespeicherte Daten noch aktuell sind? Caches sind am
effektivsten, wenn ein Browser Seiten zurückgeben kann, ohne die ursprüngliche Website kontaktieren zu müssen. Dies
birgt jedoch das Risiko, veraltete Daten zurückzugeben. Beachten Sie auch, dass die Aktualisierungsrate stark davon
abhängt, welche Daten tatsächlich zwischengespeichert sind: Während Fahrpläne für Züge kaum ändern, gilt dies nicht für
Webseiten, die aktuelle Autobahnverkehrsbedingungen oder noch schlimmer, Aktienkurse zeigen.

Was wir brauchen, ist eine Trennung zwischen Richtlinie und Mechanismus. Im Falle des Web-Cachings sollte ein Browser
beispielsweise idealerweise Einrichtungen nur zum Speichern von Dokumenten bereitstellen (d.h. ein Mechanismus) und
gleichzeitig den Benutzern ermöglichen, zu entscheiden, welche Dokumente wie lange gespeichert werden (d.h. eine
Richtlinie). In der Praxis kann dies durch die Bereitstellung eines reichen Satzes von Parametern umgesetzt werden, die
der Benutzer (dynamisch) einstellen kann. Wenn man diesen Schritt weiter geht, könnte ein Browser sogar Einrichtungen
anbieten, um Richtlinien hinzuzufügen, die ein Benutzer als separate Komponente implementiert hat.

### 1.2.4 Zuverlässigkeit

Wie der Name schon sagt, bezieht sich Zuverlässigkeit auf das Mass, inwieweit man sich darauf verlassen kann, dass ein
Computersystem wie erwartet funktioniert. Im Gegensatz zu einzelnen Computersystemen kann die Zuverlässigkeit in
verteilten Systemen aufgrund von partiellen Ausfällen recht komplex sein: Irgendwo gibt es eine fehlerhafte Komponente,
während das System insgesamt immer noch den Erwartungen zu entsprechen scheint (bis zu einem bestimmten Punkt oder
Moment). Obwohl auch Einzelcomputer-Systeme unter Ausfällen leiden können, die nicht sofort auftreten, verkompliziert
eine potenziell grosse Sammlung von vernetzten Computersystemen die Dinge erheblich. Tatsächlich sollte man davon
ausgehen, dass zu jedem Zeitpunkt immer partielle Ausfälle auftreten. Ein wichtiges Ziel von verteilten Systemen besteht
darin, diese Ausfälle zu verbergen, sowie die Wiederherstellung von diesen Ausfällen zu verbergen. Diese Verschleierung
wird als Fehlertoleranz bezeichnet.

#### Grundkonzepte

Zuverlässigkeit ist ein Begriff, der mehrere nützliche Anforderungen an verteilte Systeme abdeckt, einschliesslich der
folgenden [Kopetz und Verissimo, 1993]:

* Verfügbarkeit
* Betriebssicherheit
* Sicherheit
* Wartungsfreundlichkeit

**Verfügbarkeit** wird definiert als die Eigenschaft, dass ein System sofort einsatzbereit ist.
Allgemein bezieht es sich auf die Wahrscheinlichkeit, dass das System zu einem bestimmten Zeitpunkt korrekt arbeitet und
bereit ist, seine Funktionen im Auftrag seiner Benutzer auszuführen.
Ein hochverfügbares System ist also eines, das zu einem bestimmten Zeitpunkt höchstwahrscheinlich funktioniert.

**Zuverlässigkeit** bezieht sich auf die Eigenschaft, dass ein System kontinuierlich ohne Ausfall laufen kann.
Im Gegensatz zur Verfügbarkeit wird Zuverlässigkeit in Bezug auf ein Zeitintervall anstatt eines Augenblicks definiert.
Ein hochzuverlässiges System ist eines, das wahrscheinlich über einen relativ langen Zeitraum hinweg ohne
Unterbrechung weiterarbeiten wird.
Dies ist ein subtiler, aber wichtiger Unterschied im Vergleich zur Verfügbarkeit.
Wenn ein System durchschnittlich für eine scheinbar zufällige Millisekunde jede Stunde ausfällt, hat es eine
Verfügbarkeit von mehr als 99,9999 Prozent, ist aber dennoch unzuverlässig.
Ebenso hat ein System, das nie abstürzt, aber zwei bestimmte Wochen jeden August heruntergefahren wird, eine hohe
Zuverlässigkeit, aber nur 96 Prozent Verfügbarkeit.
Die beiden sind nicht dasselbe.

**Sicherheit** bezieht sich auf die Situation, dass bei einem vorübergehenden Versagen des Systems kein katastrophales
Ereignis eintritt.
Viele Prozesssteuerungssysteme, wie sie beispielsweise zur Steuerung von Kernkraftwerken oder zum
Senden von Personen ins All verwendet werden, müssen ein hohes Mass an Sicherheit bieten.
Wenn solche Steuerungssysteme nur für einen sehr kurzen Moment versagen, könnten die Folgen verheerend sein.
Viele Beispiele aus der Vergangenheit (und wahrscheinlich viele weitere, die noch kommen werden) zeigen, wie schwer es
ist, sichere Systeme zu bauen.

**Wartbarkeit** bezieht sich darauf, wie leicht ein ausgefallenes System repariert werden kann. Ein leicht wartbares
System kann auch eine hohe Verfügbarkeit aufweisen, insbesondere wenn Ausfälle automatisch erkannt und behoben werden
können.
Das automatische Wiederherstellen von Ausfällen ist jedoch leichter gesagt als getan.

Traditionell wurde Fehlertoleranz mit den folgenden drei Metriken in Verbindung gebracht:

* Mittlere Zeit bis zum Ausfall (MTTF): Die durchschnittliche Zeit, bis eine Komponente ausfällt.
* Mittlere Reparaturzeit (MTTR): Die durchschnittlich benötigte Zeit, um eine Komponente zu reparieren.
* Mittlere Zeit zwischen Ausfällen (MTBF): Einfach MTTF + MTTR.

Diese Metriken machen nur dann Sinn, wenn wir ein genaues Verständnis dafür haben, was ein Ausfall eigentlich ist.
Wie wir in Kapitel 8 sehen werden, ist es möglicherweise nicht so offensichtlich, das Auftreten eines Ausfalls zu
identifizieren.

#### Fehler, Störungen, Ausfälle

Ein System gilt als ausgefallen, wenn es seine Versprechen nicht erfüllen kann. Insbesondere, wenn ein verteiltes System
dazu konzipiert ist, seinen Benutzern mehrere Dienste zur Verfügung zu stellen, ist das System ausgefallen, wenn einer
oder mehrere dieser Dienste nicht (vollständig) bereitgestellt werden können. Ein Fehler ist ein Teil des
Systemzustands, der zu einem Ausfall führen kann. Beim Übertragen von Paketen über ein Netzwerk ist zu erwarten, dass
einige Pakete beschädigt sind, wenn sie beim Empfänger ankommen. Beschädigt in diesem Kontext bedeutet, dass der
Empfänger einen Bitwert möglicherweise falsch erkennt (z. B. das Lesen einer 1 anstelle einer 0) oder sogar nicht
erkennen kann, dass etwas angekommen ist.

Die Ursache eines Fehlers wird als Störung bezeichnet. Es ist klar, dass es wichtig ist, herauszufinden, was einen
Fehler verursacht hat. Eine falsche oder schlechte Übertragungsstrecke kann beispielsweise leicht dazu führen, dass
Pakete beschädigt werden. In diesem Fall ist es relativ einfach, die Störung zu beseitigen. Übertragungsfehler können
jedoch auch durch schlechte Wetterbedingungen verursacht werden, wie bei drahtlosen Netzwerken. Das Ändern des Wetters
zur Reduzierung oder Verhinderung von Fehlern ist etwas komplizierter.
Ein abgestürztes Programm ist eindeutig ein Ausfall, der möglicherweise aufgetreten ist, weil das Programm einen
Programmcode-Zweig mit einem Programmfehler betreten hat (d. h. einen Programmierfehler). Die Ursache dieses Fehlers ist
in der Regel ein Programmierer. Mit anderen Worten, der Programmierer ist die Ursache für den Fehler (Programmierbug),
der wiederum zu einem Ausfall führt (ein abgestürztes Programm).
Der Bau zuverlässiger Systeme hängt eng mit der Kontrolle von Fehlern zusammen. Wie von Avizienis et al. [2004] erklärt,
kann man zwischen Verhinderung, Tolerierung, Entfernung und Vorhersage von Fehlern unterscheiden. Für unsere Zwecke ist
das wichtigste Thema die Fehlertoleranz, was bedeutet, dass ein System seine Dienste auch bei Vorhandensein von Fehlern
erbringen kann. Beispielsweise ist es durch die Anwendung von fehlerkorrigierenden Codes zur Übertragung von Paketen
möglich, in gewissem Masse relativ schlechte Übertragungsleitungen zu tolerieren und die Wahrscheinlichkeit zu
verringern, dass ein Fehler (ein beschädigtes Paket) zu einem Ausfall führt.
Störungen werden allgemein als vorübergehend, intermittierend oder dauerhaft klassifiziert. Vorübergehende Fehler treten
einmal auf und verschwinden dann. Wenn die Operation wiederholt wird, verschwindet der Fehler. Ein Vogel, der durch den
Strahl eines Mikrowellensenders fliegt, kann zu verlorenen Bits in einem Netzwerk führen (ganz zu schweigen von einem
gerösteten Vogel). Wenn die Übertragung abläuft und erneut versucht wird, funktioniert sie wahrscheinlich beim zweiten
Mal.
Ein intermittierender Fehler tritt auf, verschwindet dann von selbst, tritt dann wieder auf und so weiter. Ein loser
Kontakt an einem Stecker verursacht oft einen intermittierenden Fehler. Intermittierende Fehler verursachen viel Ärger,
da sie schwer zu diagnostizieren sind. Normalerweise funktioniert das System gut, wenn der Fehlerbeauftragte auftaucht.
Ein dauerhafter Fehler besteht weiterhin, bis die fehlerhafte Komponente ersetzt wird. Ausgebrannte Chips,
Softwarefehler und Festplattenkopfabstürze sind Beispiele für dauerhafte Fehler.
Zuverlässige Systeme müssen auch Sicherheit bieten, insbesondere in Bezug auf Vertraulichkeit und Integrität.
Vertraulichkeit ist die Eigenschaft, dass Informationen nur autorisierten Parteien offengelegt werden, während
Integrität sicherstellt, dass Änderungen an verschiedenen Assets nur auf autorisierte Weise vorgenommen werden können.
Können wir von einem zuverlässigen System sprechen, wenn Vertraulichkeit und Integrität nicht gegeben sind? Wir kehren
zur Sicherheit zurück.

### 1.2.5 Sicherheit

Ein verteiltes System, das nicht sicher ist, ist nicht zuverlässig. Wie erwähnt, ist besondere Aufmerksamkeit
erforderlich, um Vertraulichkeit und Integrität zu gewährleisten, die beide direkt mit der autorisierten Offenlegung und
dem Zugriff auf Informationen und Ressourcen verbunden sind. In jedem Computersystem erfolgt die Autorisierung durch
Überprüfung, ob eine identifizierte Entität über die entsprechenden Zugriffsrechte verfügt. Dies bedeutet wiederum, dass
das System sicherstellen muss, dass es tatsächlich mit der richtigen Entität zu tun hat. Aus diesem Grund ist
Authentifizierung unerlässlich: Überprüfung der Korrektheit einer behaupteten Identität. Genauso wichtig ist das Konzept
des Vertrauens. Wenn ein System eine Person positiv authentifizieren kann, welchen Wert hat diese Authentifizierung,
wenn der Person nicht vertraut werden kann? Allein aus diesem Grund ist eine ordnungsgemässe Autorisierung wichtig, da
sie dazu verwendet werden kann, den Schaden zu begrenzen, den eine Person, der im Nachhinein nicht vertraut werden
konnte, verursachen kann. Zum Beispiel können in Finanzsystemen Autorisierungen den Geldbetrag begrenzen, den eine
Person zwischen verschiedenen Konten übertragen darf. Vertrauen, Authentifizierung und Autorisierung werden im Kapitel 9
ausführlich besprochen.

#### Schlüsselelemente, um Sicherheit zu verstehen

Eine wesentliche Technik, um verteilte Systeme sicher zu machen, ist die Kryptographie. Dies ist nicht der Ort in diesem
Buch, um die Kryptographie ausführlich zu besprechen (was wir auch bis Kapitel 9 verschieben), aber um zu verstehen, wie
Sicherheit in verschiedene Perspektiven in den folgenden Kapiteln passt, führen wir informell einige ihrer Grundelemente
ein.
Einfach ausgedrückt geht es bei der Sicherheit in verteilten Systemen darum, Daten mit Sicherheitsschlüsseln zu
verschlüsseln und zu entschlüsseln. Der einfachste Weg, einen Sicherheitsschlüssel K zu betrachten, besteht darin, ihn
als eine Funktion zu sehen, die auf bestimmte Daten "data" operiert. Wir verwenden die Notation K(data), um
auszudrücken, dass der Schlüssel K auf die Daten operiert.
Es gibt zwei Möglichkeiten, Daten zu verschlüsseln und zu entschlüsseln. In einem symmetrischen Kryptosystem erfolgt die
Verschlüsselung und Entschlüsselung mit einem einzigen Schlüssel. Mit E_K(data) wird die Verschlüsselung von Daten mit
dem Schlüssel EK bezeichnet, und entsprechend DK(data) für die Entschlüsselung mit dem Schlüssel DK. In einem
symmetrischen Kryptosystem wird derselbe Schlüssel sowohl für die Verschlüsselung als auch für die Entschlüsselung
verwendet, d.h.,
wenn data = DK(EK(data)) dann ist DK = EK.
Beachten Sie, dass bei einem symmetrischen Kryptosystem der Schlüssel von allen Parteien, die berechtigt sind, Daten zu
verschlüsseln oder zu entschlüsseln, geheim gehalten werden muss.

Bei einem asymmetrischen Kryptosystem sind die für
Verschlüsselung und Entschlüsselung verwendeten Schlüssel unterschiedlich. Insbesondere gibt es einen öffentlichen
Schlüssel PK, den jeder verwenden kann, und einen geheimen Schlüssel SK, der, wie der Name schon sagt, geheim gehalten
werden soll. Asymmetrische Kryptosysteme werden auch als Public-Key-Systeme bezeichnet. Verschlüsselung und
Entschlüsselung in Public-Key-Systemen können auf zwei grundsätzlich unterschiedliche Weisen verwendet werden. Erstens,
wenn Alice Daten verschlüsseln möchte, die nur von Bob entschlüsselt werden können, sollte sie Bobs öffentlichen
Schlüssel, PKB, verwenden, was zu den verschlüsselten Daten PKB(data) führt. Nur der Inhaber des zugehörigen geheimen
Schlüssels kann diese Informationen entschlüsseln, d.h. Bob, der die Operation SKB(PKB(data)) anwendet, die data
zurückgibt.

Ein zweiter, weit verbreiteter Anwendungsfall ist das Realisieren von digitalen Signaturen. Angenommen, Alice stellt
einige Daten zur Verfügung, bei denen es wichtig ist, dass eine Partei, aber lassen Sie uns annehmen, es ist Bob, sicher
wissen muss, dass sie von Alice stammen. In diesem Fall kann Alice die Daten mit ihrem geheimen Schlüssel SKA
verschlüsseln, was zu SKA(data) führt. Wenn sichergestellt werden kann, dass der zugehörige öffentliche Schlüssel PKA
tatsächlich Alice gehört, dann ist das erfolgreiche Entschlüsseln von SKA(data) zu data der Beweis, dass Alice von data
weiss: Sie ist die einzige, die den geheimen Schlüssel SKA besitzt. Natürlich müssen wir die Annahme treffen, dass Alice
tatsächlich die einzige ist, die SKA besitzt. Zu einigen dieser Annahmen kehren wir in Kapitel 9 zurück.
Wie sich herausstellt, kommt der Beweis, dass eine Entität einige Daten gesehen hat oder von ihnen weiss, häufig in
gesicherten verteilten Systemen vor. Die praktische Platzierung von digitalen Signaturen erfolgt in der Regel
effizienter durch eine Hash-Funktion. Eine Hash-Funktion H hat die Eigenschaft, dass sie bei der Operation auf einige
Daten, d.h. H(data), eine Zeichenkette fester Länge zurückgibt, unabhängig von der Länge der Daten. Jede Änderung von
data zu data* führt zu einem anderen Hash-Wert H(data*). Darüber hinaus ist es in der Praxis rechnerisch unmöglich, die
ursprünglichen Daten anhand eines Hash-Wertes h zu entdecken. Was das alles bedeutet, ist, dass Alice für das Anbringen
einer digitalen Signatur sig = SKA(H(data)) als ihre Signatur berechnet und Bob über data, H und sig informiert. Bob
kann dann diese Signatur überprüfen, indem er PKA(sig) berechnet und überprüft, ob sie mit dem Wert H(data)
übereinstimmt.

#### Verwendung von Kryptographie in verteilten Systemen

Die Anwendung von Kryptographie in verteilten Systemen hat viele Formen. Neben ihrer allgemeinen Verwendung für
Verschlüsselung und Entschlüsselung sind insbesondere Authentifizierungsprotokolle, in denen ein System überprüfen kann,
ob es tatsächlich mit der richtigen Entität zu tun hat, und Verteilungsprotokolle, mit denen Schlüssel oder vertrauliche
Informationen sicher zwischen Parteien ausgetauscht werden können, von Bedeutung. Schliesslich gibt es noch eine Klasse
von Protokollen, die versuchen, festzustellen, ob eine Partei ehrlich ist oder ob sie versucht, das System in
irgendeiner Weise zu täuschen.
In den folgenden Kapiteln werden wir sehen, wie diese verschiedenen Protokolle und Techniken in verteilten Systemen
eingesetzt werden können, um die erforderliche Sicherheit zu gewährleisten. Was jedoch über alle diese Protokolle hinweg
klar sein muss, ist, dass es, wie in jedem System, eine Kosten-Nutzen-Analyse gibt. Sicherheit ist nicht umsonst, und
wir müssen bereit sein, dafür zu zahlen, sei es in Form von zusätzlichen Hardwarekomponenten, zusätzlicher Software,
zusätzlichen Kommunikationskosten oder in Form von weniger Benutzerfreundlichkeit. Es ist wichtig, einen Kompromiss
zwischen der gewünschten Sicherheit und den damit verbundenen Kosten zu finden.

### 1.2.6 Skalierbarkeit

Für viele von uns ist die weltweite Vernetzung durch das Internet genauso alltäglich wie die Möglichkeit, ein Paket an
jeden Ort der Welt zu senden. Früher waren wir es gewohnt, relativ leistungsstarke Desktop-Computer für Büroanwendungen
und Speicherung zu haben. Heutzutage werden solche Anwendungen und Dienste jedoch in das verschoben, was als "die Cloud"
bezeichnet wird. Dies führt zu einem Anstieg von viel kleineren vernetzten Geräten wie Tablet-Computern oder sogar
reinen Cloud-Laptops wie Googles Chromebook. Vor diesem Hintergrund ist die Skalierbarkeit zu einem der wichtigsten
Designziele für Entwickler verteilter Systeme geworden.

#### Skalierbarkeitsdimensionen

Die Skalierbarkeit eines Systems kann entlang von mindestens drei verschiedenen Dimensionen gemessen werden (
siehe [Neuman, 1994]):

* **Grössenskalierbarkeit:** Ein System kann bezüglich seiner Grösse skalierbar sein,
  was bedeutet, dass wir dem System problemlos weitere Benutzer und Ressourcen hinzufügen können, ohne dass ein
  spürbarer Leistungsverlust auftritt.
* **Geografische Skalierbarkeit:** Ein geografisch skalierbares System ist eines, bei dem Benutzer und Ressourcen weit
  voneinander entfernt liegen können, aber Kommunikationsverzögerungen kaum bemerkt werden.
* **Administrative Skalierbarkeit:** Ein administrativ skalierbares System ist eines, das auch dann noch leicht
  verwaltet
  werden kann, wenn es viele unabhängige Verwaltungsorganisationen umfasst.

Lassen Sie uns einen genaueren Blick auf jede dieser drei Skalierbarkeitsdimensionen werfen.

**Grössenskalierbarkeit**: Wenn ein System skalieren muss, müssen sehr unterschiedliche Arten von Problemen gelöst
werden. Betrachten wir zunächst die Skalierung hinsichtlich der Grösse. Wenn mehr Benutzer oder Ressourcen unterstützt
werden müssen, stossen wir oft auf die Grenzen zentralisierter Dienste, obwohl oft aus sehr unterschiedlichen Gründen.
Viele Dienste sind in dem Sinne zentralisiert, dass sie von einem einzelnen Server auf einer spezifischen Maschine im
verteilten System implementiert werden. In einem moderneren Setting könnten wir eine Gruppe von zusammenarbeitenden
Servern haben, die auf einem Cluster von eng gekoppelten Maschinen an einem Ort koexistieren. Das Problem mit diesem
Schema ist offensichtlich: Der Server oder die Gruppe von Servern kann einfach zum Engpass werden, wenn er eine
wachsende Anzahl von Anfragen bearbeiten muss. Um zu veranschaulichen, wie dies passieren kann, nehmen wir an, dass ein
Dienst auf einer einzelnen Maschine implementiert ist. Es gibt im Wesentlichen drei Hauptursachen für Engpässe:

- Die Rechenkapazität, begrenzt durch die CPUs
- Die Speicherkapazität, einschliesslich der I/O-Übertragungsrate
- Das Netzwerk zwischen dem Benutzer und dem zentralisierten Dienst

Betrachten wir zunächst die Rechenkapazität. Stellen Sie sich einen Dienst vor, der optimale Routen unter
Berücksichtigung von Echtzeit-Verkehrsinformationen berechnet. Es ist nicht schwer vorstellbar, dass dies
hauptsächlich ein rechenintensiver Dienst sein könnte, der mehrere (zehn) Sekunden benötigt, um eine Anfrage
abzuschliessen. Wenn nur eine einzige Maschine verfügbar ist, wird selbst ein modernes Hochleistungssystem schliesslich
auf Probleme stossen, wenn die Anzahl der Anfragen über einen bestimmten Punkt hinaus ansteigt.

Ebenso werden wir, aber aus anderen Gründen, auf Probleme stossen, wenn wir einen hauptsächlich I/O-gebundenen Dienst
haben. Ein typisches Beispiel ist eine schlecht gestaltete zentralisierte Suchmaschine. Das Problem mit inhaltsbasierten
Suchanfragen besteht darin, dass wir im Wesentlichen eine Abfrage gegen einen gesamten Datensatz abgleichen müssen.
Selbst mit fortgeschrittenen Indexierungstechniken könnten wir immer noch vor dem Problem stehen, eine riesige
Datenmenge verarbeiten zu müssen, die die Haupt-Speicherkapazität der Maschine, auf der der Dienst läuft, übersteigt.
Als Konsequenz wird ein Grossteil der Verarbeitungszeit durch die vergleichsweise langsamen Festplattenzugriffe und den
Datentransfer zwischen Festplatte und Hauptspeicher bestimmt. Einfach mehr oder schnellere Festplatten hinzuzufügen,
wird sich als keine nachhaltige Lösung erweisen, wenn die Anzahl der Anfragen weiter zunimmt.

Schliesslich kann auch das Netzwerk zwischen dem Benutzer und dem Dienst die Ursache für schlechte Skalierbarkeit sein.
Stellen Sie sich einen Video-on-Demand-Dienst vor, der hochwertige Videos an mehrere Benutzer streamen muss. Ein
Videostream kann leicht eine Bandbreite von 8 bis 10 Mbps benötigen. Wenn ein Dienst Punkt-zu-Punkt-Verbindungen mit
seinen Kunden einrichtet, kann er bald an die Grenzen der Netzwerkkapazität seiner eigenen ausgehenden
Übertragungsleitungen stossen.

Es gibt mehrere Lösungen, um die Grössenskalierbarkeit zu bewältigen, die wir weiter unten diskutieren werden, nachdem
wir geografische und administrative Skalierbarkeit betrachtet haben.

**Geografische Skalierbarkeit**: Ein weiteres Problem, das die geografische Skalierbarkeit beeinträchtigt, ist, dass
viele von ihnen auf synchroner Kommunikation basieren. Bei dieser Form der Kommunikation blockiert eine Partei, die
einen Dienst anfordert (in der Regel als Client bezeichnet), bis eine Antwort vom den Dienst implementierenden Server
zurückgesendet wird. Genauer gesagt, sehen wir oft ein Kommunikationsmuster, das aus vielen Client-Server-Interaktionen
besteht, wie es bei Datenbanktransaktionen der Fall sein kann. Dieser Ansatz funktioniert in der Regel gut in LANs, bei
denen die Kommunikation zwischen zwei Maschinen oft höchstens einige hundert Mikrosekunden beträgt. In einem
Wide-Area-System müssen wir jedoch berücksichtigen, dass die Interprozesskommunikation möglicherweise hunderte von
Millisekunden dauert, das ist drei Grössenordnungen langsamer. Das Erstellen von Anwendungen unter Verwendung synchroner
Kommunikation in Wide-Area-Systemen erfordert viel Vorsicht (aber es ist nicht unmöglich). Wir werden auch darauf später
eingehen.

**Administrative Skalierbarkeit:** Schliesslich ist eine schwierige und oft offene Frage, wie man ein verteiltes System
über
mehrere, unabhängige administrative Domänen hinweg skalieren kann. Ein Hauptproblem, das gelöst werden muss, ist das der
konfligierenden Richtlinien bezüglich Ressourcennutzung (und Bezahlung), Verwaltung und Sicherheit. Zur
Veranschaulichung haben Wissenschaftler viele Jahre nach Lösungen gesucht, um ihre (oft teure) Ausrüstung in dem, was
als Rechen-Grid bekannt ist, zu teilen. In diesen Grids wird ein globales dezentralisiertes System als Föderation
lokaler verteilter Systeme aufgebaut, wodurch ein Programm, das auf einem Computer bei einer Organisation A läuft,
direkt auf Ressourcen bei der Organisation B zugreifen kann.

Viele Komponenten eines verteilten Systems, die innerhalb einer einzigen Domäne liegen, können oft von Benutzern, die
innerhalb derselben Domäne operieren, als vertrauenswürdig angesehen werden. In solchen Fällen haben die
Systemadministratoren möglicherweise Anwendungen getestet und zertifiziert und besondere Massnahmen ergriffen, um
sicherzustellen, dass solche Komponenten nicht manipuliert werden können. Im Wesentlichen vertrauen die Benutzer ihren
Systemadministratoren. Dieses Vertrauen erweitert sich jedoch nicht natürlich über die Grenzen von Domänen hinweg.
Wenn sich ein verteiltes System auf eine andere Domäne ausdehnt, müssen zwei Arten von Sicherheitsmassnahmen ergriffen
werden. Erstens muss sich das verteilte System gegen bösartige Angriffe aus der neuen Domäne schützen. Zum Beispiel
haben Benutzer aus der neuen Domäne möglicherweise nur Lesezugriff auf das Dateisystem in seiner ursprünglichen Domäne.
Ebenso dürfen Einrichtungen wie teure Bildsatzsysteme oder Hochleistungscomputer nicht für nicht autorisierte Benutzer
verfügbar gemacht werden. Zweitens muss sich die neue Domäne gegen bösartige Angriffe aus dem verteilten System
schützen. Ein typisches Beispiel ist das Herunterladen von Programmen, wie im Fall des föderierten Lernens.
Grundsätzlich weiss die neue Domäne nicht, was sie von solch fremdem Code erwarten kann. Das Problem, wie wir in Kapitel
9 sehen werden, besteht darin, diese Beschränkungen durchzusetzen.
Als Gegenbeispiel für verteilte Systeme, die mehrere administrative Domänen umspannen und anscheinend keine Probleme mit
der administrativen Skalierbarkeit haben, betrachten Sie moderne File-Sharing-Peer-to-Peer-Netzwerke. In diesen Fällen
installieren Endbenutzer einfach ein Programm, das verteilte Such- und Downloadfunktionen implementiert, und können
innerhalb von Minuten mit dem Herunterladen von Dateien beginnen. Andere Beispiele sind Peer-to-Peer-Anwendungen für
Internettelefonie wie ältere Skype-Systeme [Baset und Schulzrinne, 2006] und (wieder ältere) Peer-unterstützte
Audio-Streaming-Anwendungen wie Spotify [Kreitz und Niemelä, 2010]. Eine modernere Anwendung (die ihre Skalierbarkeit
noch beweisen muss) sind Blockchains. Was diese dezentralisierten Systeme gemeinsam haben, ist, dass Endbenutzer, und
nicht administrative Entitäten, zusammenarbeiten, um das System am Laufen zu halten. Im besten Fall können zugrunde
liegende administrative Organisationen wie Internetdienstanbieter (ISPs) den Netzwerkverkehr überwachen, den diese
Peer-to-Peer-Systeme verursachen.

#### Skalierungstechniken

Nachdem wir einige Skalierbarkeitsprobleme besprochen haben, stellt sich uns die Frage, wie diese Probleme generell
gelöst werden können. In den meisten Fällen treten Skalierbarkeitsprobleme in verteilten Systemen als Leistungsprobleme
auf, die durch die begrenzte Kapazität von Servern und Netzwerken verursacht werden. Eine einfache Erhöhung ihrer
Kapazität (z.B. durch Erhöhen des Speichers, Aufrüsten der CPUs oder Ersetzen von Netzwerkmodulen) ist oft eine Lösung,
die als "vertikale Skalierung" bezeichnet wird. Bei der sogenannten "horizontalen Skalierung", also um die Erweiterung
des verteilten
Systems durch mehr Maschinen, gibt es im Grunde nur drei Techniken, die wir anwenden können:
Kommunikationslatenzen verbergen, Arbeit verteilen und Replikation (siehe auch Neuman [1994]).

#### Kommunikationslatenzen verbergen

Das Verbergen von Kommunikationslatenzen ist im Fall der geografischen Skalierbarkeit anwendbar. Die Grundidee ist
einfach: Versuchen Sie, so wenig wie möglich auf Antworten auf Anfragen zu warten. Wenn ein Service auf einer
enfernten Maschine angefordert wurde, besteht eine Alternative zum Warten auf eine Antwort vom Server darin, andere
nützliche Arbeiten auf der Anfordererseite durchzuführen. Im Wesentlichen bedeutet dies, die anfordernde Anwendung so zu
konstruieren, dass sie nur asynchrone Kommunikation verwendet. Wenn eine Antwort eintrifft, wird die Anwendung
unterbrochen und ein spezieller Handler aufgerufen, um die zuvor ausgestellte Anfrage abzuschliessen. Asynchrone
Kommunikation kann oft in Batch-Verarbeitungssystemen und parallelen Anwendungen verwendet werden, bei denen unabhängige
Aufgaben zur Ausführung geplant werden können, während eine andere Aufgabe auf die Fertigstellung der Kommunikation
wartet. Alternativ kann ein neuer Thread gestartet werden, um die Anfrage auszuführen. Obwohl er beim Warten
auf die Antwort blockiert wird, können andere Threads im Prozess weiterlaufen.

Es gibt jedoch viele Anwendungen, die asynchrone Kommunikation nicht effektiv nutzen können. Zum Beispiel haben
interaktive Anwendungen, bei denen ein Benutzer eine Anfrage sendet, im Allgemeinen nichts Besseres zu tun, als auf die
Antwort zu warten. In solchen Fällen ist es viel besser, die gesamte Kommunikation zu reduzieren, zum Beispiel, indem
ein Teil der Berechnung, die normalerweise auf dem Server durchgeführt wird, zum Client-Prozess verschoben wird, der den
Service anfordert. Ein typisches Beispiel, bei dem dieser Ansatz funktioniert, ist der Zugriff auf Datenbanken mithilfe
von Formularen. Das Ausfüllen von Formularen kann durch das Senden einer separaten Nachricht für jedes Feld und das
Warten auf eine Bestätigung vom Server erfolgen, wie in Abbildung 1.6(a) gezeigt. Zum Beispiel kann der Server auf
syntaktische Fehler überprüfen, bevor er einen Eintrag akzeptiert.

Eine viel bessere Lösung besteht darin, den Code zum Ausfüllen des Formulars und möglicherweise das Überprüfen der
Einträge zum Client zu senden und den Client ein ausgefülltes Formular zurückgeben zu lassen, wie in Abbildung 1.6(b)
gezeigt. Dieser Ansatz des Code-Versands wird vom Web durch JavaScript weitgehend unterstützt.

#### Partitionierung und Verteilung

Eine weitere wichtige Skalierungstechnik ist die Partitionierung und Verteilung, die darin besteht, eine Komponente oder
andere Ressource in kleinere Teile zu zerlegen und diese Teile anschliessend im System zu verteilen. Ein gutes Beispiel
für Partitionierung und Verteilung ist das Internet Domain Name System (DNS). Der DNS-Namensraum ist hierarchisch in
einen Baum von Domänen organisiert, der in nicht überlappende Zonen unterteilt ist, wie für das ursprüngliche DNS in
Abbildung 1.7 gezeigt. Die Namen in jeder Zone werden von einem einzelnen Namensserver gehandhabt. Ohne jetzt zu sehr
ins Detail zu gehen (wir kehren in Kapitel 6 ausführlich zu DNS zurück), kann man sich jeden Pfadnamen als den Namen
eines Hosts im Internet vorstellen und ist daher mit einer Netzwerkadresse dieses Hosts verknüpft. Grundsätzlich
bedeutet die Auflösung eines Namens, die Netzwerkadresse des zugehörigen Hosts zurückzugeben. Betrachten Sie zum
Beispiel den Namen flits.cs.vu.nl. Um diesen Namen aufzulösen, wird er zuerst an den Server von Zone Z1 (siehe Abbildung
1.7) übergeben, der die Adresse des Servers für Zone Z2 zurückgibt, an den der Rest des Namens, flits.cs.vu, übergeben
werden kann. Der Server für Z2 gibt die Adresse des Servers für Zone Z3 zurück, der den letzten Teil des Namens
verarbeiten kann, und gibt die Adresse des zugehörigen Hosts zurück.

Diese Beispiele veranschaulichen, wie der von DNS bereitgestellte Namensdienst über mehrere Maschinen verteilt ist und
so vermieden wird, dass ein einzelner Server alle Anfragen zur Namensauflösung bearbeiten muss.

Als weiteres Beispiel betrachten Sie das World Wide Web. Für die meisten Benutzer scheint das Web ein riesiges
dokumentenbasiertes Informationssystem zu sein, bei dem jedes Dokument seinen eigenen eindeutigen Namen in Form einer
URL hat. Konzeptionell könnte es sogar so aussehen, als gäbe es nur einen einzigen Server. Das Web ist jedoch physisch
partitioniert und auf einige hundert Millionen Server verteilt, von denen jeder oft mehrere Websites oder Teile von
Websites verwaltet. Der Name des Servers, der ein Dokument verwaltet, ist in die URL dieses Dokuments kodiert. Nur
aufgrund dieser Verteilung von Dokumenten konnte das Web auf seine jetzige Grösse skalieren. Beachten Sie jedoch, dass es
fast unmöglich ist herauszufinden, wie viele Server Webdienste anbieten: Eine Website ist heute viel mehr als nur einige
statische Webdokumente.

#### Replikation

In Anbetracht der Tatsache, dass Skalierbarkeitsprobleme oft in Form von Leistungsabfall auftreten, ist es im
Allgemeinen eine gute Idee, Komponenten oder Ressourcen usw. in einem verteilten System zu replizieren.
Replikation erhöht nicht nur die Verfügbarkeit, sondern hilft auch, die Last zwischen den Komponenten auszugleichen, was
zu einer besseren Leistung führt. Darüber hinaus kann in geografisch weit verteilten Systemen eine in der Nähe
befindliche Kopie viele der zuvor erwähnten Kommunikationslatenzprobleme verbergen.

Caching ist eine spezielle Form der Replikation, obwohl die Unterscheidung zwischen den beiden oft schwer zu treffen
oder sogar künstlich ist. Wie im Falle der Replikation führt Caching dazu, dass eine Kopie einer Ressource erstellt
wird, im Allgemeinen in der Nähe des Kunden, der auf diese Ressource zugreift. Im Gegensatz zur Replikation ist Caching
jedoch eine Entscheidung, die vom Kunden einer Ressource und nicht vom Besitzer einer Ressource getroffen wird.

Es gibt jedoch einen ernsthaften Nachteil beim Caching und bei der Replikation, der die Skalierbarkeit negativ
beeinflussen kann. Da wir nun mehrere Kopien einer Ressource haben, macht das Ändern einer Kopie diese anders als die
anderen. Folglich führen Caching und Replikation zu Konsistenzproblemen.

Inwieweit Inkonsistenzen toleriert werden können, hängt von der Nutzung einer Ressource ab. Zum Beispiel finden es viele
Webnutzer akzeptabel, dass ihr Browser ein zwischengespeichertes Dokument zurückgibt, dessen Gültigkeit in den letzten
Minuten nicht überprüft wurde. Es gibt jedoch auch viele Fälle, in denen starke Konsistenzgarantien erfüllt werden
müssen, wie im Fall von elektronischen Börsen und Auktionen. Das Problem mit starker Konsistenz ist, dass eine
Aktualisierung sofort zu allen anderen Kopien weitergeleitet werden muss. Ausserdem ist es oft auch erforderlich, dass
Updates gleichzeitig erfolgen, wenn zwei Aktualisierungen gleichzeitig stattfinden, was ein globales Ordnungsproblem
einführt. Um die Dinge noch schlimmer zu machen, kann die Kombination von Konsistenz mit wünschenswerten Eigenschaften
wie Verfügbarkeit einfach unmöglich sein, wie wir in Kapitel 8 diskutieren.

Replikation erfordert daher oft einen globalen Synchronisationsmechanismus. Leider sind solche Mechanismen extrem schwer
oder sogar unmöglich in einer skalierbaren Weise zu implementieren, schon allein weil Netzwerklatenzen eine natürliche
untere Grenze haben. Folglich kann die Skalierung durch Replikation andere, inhärent nicht skalierbare Lösungen
einführen. Wir kehren in Kapitel 7 ausführlich zu Replikation und Konsistenz zurück.

#### Diskussion

Beim Betrachten dieser Skalierungstechniken könnte man argumentieren, dass die Grössenskalierbarkeit aus technischer
Sicht das geringste Problem darstellt. Oft kann die Erhöhung der Kapazität einer Maschine Abhilfe schaffen, obwohl
möglicherweise hohe monetäre Kosten anfallen. Geografische Skalierbarkeit ist ein weitaus schwierigeres Problem, da
Netzwerklatenzen natürlich von unten begrenzt sind. Infolgedessen könnten wir gezwungen sein, Daten an Orte zu kopieren,
die den Kunden nahe liegen, was zu Problemen bei der Aufrechterhaltung konsistenter Kopien führt. Die Praxis zeigt, dass
die Kombination von Verteilungs-, Replikations- und Caching-Techniken mit verschiedenen Formen der Konsistenz im
Allgemeinen zu akzeptablen Lösungen führt. Schliesslich scheint die administrative Skalierbarkeit das schwierigste zu
lösende Problem zu sein, teilweise weil wir uns mit nichttechnischen Fragen auseinandersetzen müssen, wie der Politik
von Organisationen und der menschlichen Zusammenarbeit. Die Einführung und nun weit verbreitete Nutzung von
Peer-to-Peer-Technologie hat erfolgreich demonstriert, was erreicht werden kann, wenn Endbenutzer die Kontrolle
übernehmen [Lua et al., 2005; Oram, 2001]. Allerdings sind Peer-to-Peer-Netzwerke offensichtlich nicht die universelle
Lösung für alle Probleme der administrativen Skalierbarkeit.

## 1.4 Fallstricke

Es sollte mittlerweile klar sein, dass die Entwicklung eines verteilten Systems eine gewaltige Aufgabe ist.
Wie wir in diesem Buch mehrmals sehen werden, gibt es so viele Probleme zu berücksichtigen, dass das Ergebnis nur
Komplexität zu sein scheint.
Dennoch können durch die Befolgung mehrerer Entwurfsprinzipien verteilte Systeme entwickelt werden, die den in diesem
Kapitel festgelegten Zielen gut entsprechen.
Verteilte Systeme unterscheiden sich von traditioneller Software, da Komponenten über ein Netzwerk verteilt sind.
Wenn diese Dispersion während der Entwurfszeit nicht berücksichtigt wird, werden viele Systeme unnötig komplex und
werden Fehler gemacht, die später behoben werden müssen.
Peter Deutsch, damals bei Sun Microsystems tätig, formulierte folgende falschen Annahmen, die viele bei der Entwicklung
einer verteilten Anwendung zum ersten Mal treffen:

- Das Netzwerk ist zuverlässig
- Das Netzwerk ist sicher
- Das Netzwerk ist homogen
- Die Topologie ändert sich nicht
- Latenz ist null
- Bandbreite ist unendlich
- Transportkosten sind null
- Es gibt einen Administrator

Beachten Sie, wie diese Annahmen sich auf Eigenschaften beziehen, die einzigartig für verteilte Systeme sind:
Zuverlässigkeit, Sicherheit, Heterogenität und Netzwerk-Topologie; Latenz und Bandbreite; Transportkosten; und
schliesslich administrative Bereiche.
Bei der Entwicklung von nicht verteilten Anwendungen werden die meisten dieser Probleme wahrscheinlich nicht auftreten.

Die meisten in diesem Buch besprochenen Prinzipien beziehen sich direkt auf diese Annahmen.
In allen Fällen diskutieren wir Lösungen für Probleme, die durch die Tatsache verursacht werden, dass eine oder mehrere
Annahmen falsch sind.
Zum Beispiel gibt es einfach keine zuverlässigen Netzwerke, was zur Unmöglichkeit der Erreichung von Ausfallstransparenz
führt.
Wir widmen ein ganzes Kapitel der Tatsache, dass netzwerkgestützte Kommunikation von Natur aus unsicher ist.
Wir haben bereits argumentiert, dass verteilte Systeme offen sein müssen und die Heterogenität berücksichtigen sollten.
Ebenso behandeln wir beim Diskutieren von Replikation zur Lösung von Skalierbarkeitsproblemen im Wesentlichen Latenz-
und Bandbreitenprobleme.
Wir werden auch an verschiedenen Stellen in diesem Buch auf Managementprobleme eingehen.

## 1.5 Zusammenfassung

Ein verteiltes System ist eine Sammlung von vernetzten Computersystemen, in denen Prozesse und Ressourcen auf
verschiedene Computer verteilt sind. Wir unterscheiden zwischen ausreichend und notwendigerweise verteilt, wobei
letzteres sich auf dezentrale Systeme bezieht. Diese Unterscheidung ist wichtig, da das Verteilen von Prozessen und
Ressourcen nicht als eigenständiges Ziel betrachtet werden kann. Stattdessen resultieren die meisten Entscheidungen für
ein verteiltes System aus dem Bedarf, die Leistung eines einzelnen Computersystems hinsichtlich Zuverlässigkeit,
Skalierbarkeit und Effizienz zu verbessern. Aber wenn man bedenkt, dass die meisten zentralisierten Systeme immer noch
viel einfacher zu verwalten und zu warten sind, sollte man sich überlegen, Prozesse und Ressourcen zu verteilen. Es gibt
auch Fälle, in denen es einfach keine Wahl gibt, zum Beispiel beim Verbinden von Systemen verschiedener Organisationen
oder wenn Computer einfach aus verschiedenen Standorten arbeiten (wie beim mobilen Computing).
Entwurfsziele für verteilte Systeme beinhalten das Teilen von Ressourcen und das Sicherstellen von Offenheit. Immer
wichtiger wird das Entwerfen von sicheren verteilten Systemen. Darüber hinaus zielen Designer darauf ab, viele der
Feinheiten im Zusammenhang mit der Verteilung von Prozessen, Daten und Kontrolle zu verbergen. Diese
Verteilungstransparenz kommt jedoch nicht nur mit einem Leistungspreis, in der Praxis kann sie nie vollständig erreicht
werden. Die Tatsache, dass Abwägungen zwischen verschiedenen Formen der Verteilungstransparenz getroffen werden müssen,
ist inhärent im Design von verteilten Systemen und kann deren Verständnis leicht verkomplizieren. Ein spezifisches,
schwieriges Entwurfsziel, das nicht immer gut mit der Erreichung von Verteilungstransparenz harmoniert, ist die
Skalierbarkeit. Dies gilt insbesondere für geografische Skalierbarkeit, bei der das Verbergen von Latenzen und
Bandbreitenbeschränkungen schwierig sein kann. Ebenso kann die administrative Skalierbarkeit, bei der ein System so
gestaltet ist, dass es mehrere administrative Bereiche umfasst, leicht mit Zielen für die Erreichung von
Verteilungstransparenz in Konflikt geraten.
Es gibt verschiedene Typen von verteilten Systemen, die als auf die Unterstützung von Berechnungen,
Informationsverarbeitung und Durchdringung ausgerichtet klassifiziert werden können. Verteilte Computing-Systeme werden
typischerweise für Hochleistungsanwendungen eingesetzt, die oft aus dem Bereich des parallelen Computings stammen. Ein
Bereich, der aus der parallelen Verarbeitung entstand, war anfangs das Grid-Computing mit einem starken Fokus auf das
weltweite Teilen von Ressourcen, was schliesslich zu dem führte, was heute als Cloud-Computing bekannt ist.
Cloud-Computing geht über das Hochleistungscomputing hinaus und unterstützt auch verteilte Systeme in traditionellen
Büroumgebungen, in denen Datenbanken eine wichtige Rolle spielen. Typischerweise werden in diesen Umgebungen
Transaktionsverarbeitungssysteme eingesetzt. Schliesslich ist eine aufkommende Klasse von verteilten Systemen dort, wo
die Komponenten klein sind, das System ad hoc zusammengestellt wird, vor allem aber nicht mehr durch einen
Systemadministrator verwaltet wird. Diese letzte Klasse wird typischerweise von durchdringenden Computingsystemen,
einschliesslich mobiler Computing-Systeme sowie sensorreichen Umgebungen, dargestellt.
Die Angelegenheiten werden weiter kompliziert durch die Tatsache, dass viele Entwickler ursprünglich Annahmen über das
zugrundeliegende Netzwerk treffen, die grundlegend falsch sind. Später, wenn die Annahmen fallen gelassen werden, kann
es schwierig sein, unerwünschtes Verhalten zu maskieren. Ein typisches Beispiel ist die Annahme, dass Netzwerk-Latenz
nicht signifikant ist. Andere Fallstricke sind die Annahme, dass das Netzwerk zuverlässig, statisch, sicher und homogen
ist.

# 2 Architekturen

Verteilte Systeme sind oft komplexe Software, deren Komponenten per Definition auf mehrere Maschinen verteilt
sind. Um ihre Komplexität zu beherrschen, ist es entscheidend, dass diese Systeme sinnvoll organisiert werden. Es
gibt verschiedene Ansichten darüber, wie die Organisation eines verteilten Systems betrachtet werden kann, aber eine
offensichtliche ist die Unterscheidung zwischen der logischen Organisation der Softwarekomponenten
einerseits und der tatsächlichen physischen Realisierung andererseits.
Die Organisation verteilter Systeme dreht sich meist um die Softwarekomponenten, die das System bilden. Diese
Softwarearchitekturen sagen uns, wie die verschiedenen Softwarekomponenten organisiert sein sollten und wie sie
interagieren sollten. In diesem Kapitel werden wir zunächst einigen häufig angewandten architektonischen Stilen zur
Organisation von (verteilten) Computersystemen Beachtung schenken.
Ein wichtiges Ziel verteilter Systeme ist es, Anwendungen von zugrunde liegenden Plattformen zu trennen, indem sie eine
sogenannte Middleware-Schicht bereitstellen. Die Verwendung einer solchen Schicht ist eine wichtige architektonische
Entscheidung, und ihr Hauptzweck ist es, Verteilungstransparenz zu bieten. Allerdings müssen Kompromisse eingegangen
werden, um Transparenz zu erreichen, was zu verschiedenen Techniken geführt hat, um die Middleware an die Bedürfnisse
der Anwendungen anzupassen, die sie nutzen. Wir diskutieren einige der häufiger angewandten Techniken, da sie die
Organisation der Middleware selbst beeinflussen.
Die tatsächliche Realisierung eines verteilten Systems erfordert, dass wir Softwarekomponenten auf echten Maschinen
instanziieren und platzieren. Dabei können viele Entscheidungen getroffen werden. Die endgültige Instanzierung einer
Softwarearchitektur wird auch als Systemarchitektur bezeichnet. In diesem Kapitel werden wir uns traditionellen
zentralisierten Architekturen zuwenden, in denen ein einziger Server die meisten Softwarekomponenten (und somit
Funktionalitäten) implementiert, während entfernte Clients diesen Server mit einfachen Kommunikationsmitteln zugreifen
können. Zusätzlich betrachten wir dezentrale Peer-to-Peer-Architekturen, in denen alle Knoten mehr oder weniger
gleichberechtigte Rollen spielen. Viele reale verteilte Systeme sind oft in einer hybriden Art und Weise organisiert,
indem sie Elemente aus zentralisierten und dezentralisierten Architekturen kombinieren. Wir diskutieren mehrere
Beispiele, die die Komplexität vieler realer verteilter Systeme veranschaulichen.

## 2.1 Architekturstile

Wir beginnen unsere Diskussion über Architekturen, indem wir zunächst die logische Organisation eines verteilten Systems
in Softwarekomponenten betrachten, auch als seine Softwarearchitektur 
bezeichnet [Bass et al., 2021; Richards und Ford, 2020]. Die Forschung zu Softwarearchitekturen hat sich erheblich
weiterentwickelt, und es wird mittlerweile allgemein akzeptiert, dass das Entwerfen oder Übernehmen einer Architektur
entscheidend für die erfolgreiche Entwicklung grosser Softwaresysteme ist.

Für unsere Diskussion ist der Begriff des Architekturstils wichtig. Ein solcher Stil wird in Bezug auf Komponenten
formuliert, die Art und Weise, wie Komponenten miteinander verbunden sind, die zwischen Komponenten ausgetauschten Daten
und schliesslich, wie diese Elemente gemeinsam zu einem System konfiguriert werden. Eine Komponente ist eine modulare
Einheit mit wohldefinierten erforderlichen und bereitgestellten Schnittstellen, die innerhalb ihrer Umgebung
austauschbar ist [OMG, 2004]. Dass eine Komponente ersetzt werden kann, insbesondere während ein System weiter betrieben
wird, ist wichtig. Denn oft ist es keine Option, ein System für Wartungsarbeiten herunterzufahren. Im besten Fall können
nur Teile davon vorübergehend ausser Betrieb gesetzt werden. Das Ersetzen einer Komponente ist nur möglich, wenn ihre
Schnittstellen unberührt bleiben. In der Praxis sehen wir, dass das Ersetzen oder Aktualisieren einer Komponente
bedeutet, dass ein Teil eines Systems (wie ein Server) ein reguläres Update durchführt und auf die aktualisierten
Komponenten umschaltet, sobald deren Installation abgeschlossen ist. Besondere Massnahmen können erforderlich sein, wenn
ein Teil des verteilten Systems neu gestartet werden muss, damit die Aktualisierungen wirksam werden. Solche Massnahmen
können das Vorhandensein replizierter Standbys umfassen, die die Funktion übernehmen, während der teilweise Neustart
stattfindet.

Ein etwas schwieriger zu erfassendes Konzept ist das eines Konnektors, der allgemein als ein Mechanismus beschrieben
wird, der Kommunikation, Koordination oder Kooperation zwischen Komponenten vermittelt [Bass et al., 2021]. Ein
Konnektor kann beispielsweise durch Einrichtungen für (ferngesteuerte) Prozeduraufrufe, Nachrichtenübermittlung oder
Datenstreaming gebildet werden. Mit anderen Worten, ein Konnektor ermöglicht den Fluss von Steuerung und Daten zwischen
Komponenten.

Mit Komponenten und Konnektoren können wir zu verschiedenen Konfigurationen gelangen, die wiederum in Architekturstile
klassifiziert wurden. Mehrere Stile wurden mittlerweile identifiziert, von denen die wichtigsten für verteilte Systeme
sind:

- Schichtenarchitekturen
- Serviceorientierte-Architekturen
- Publish-Subscribe-Architekturen

Im Folgenden diskutieren wir jeden dieser Stile separat. Wir merken im Voraus an, dass in den meisten realen verteilten
Systemen viele Stile kombiniert werden. Insbesondere das Verfolgen eines Ansatzes, bei dem ein System in mehrere (
logische) Schichten unterteilt wird, ist ein so universelles Prinzip, dass es im Allgemeinen mit den meisten anderen
Architekturstilen kombiniert wird.

### 2.1.1 Schichtenarchitekturen

Die grundlegende Idee des geschichteten Stils ist einfach: Komponenten werden in einer geschichteten Weise organisiert,
wobei eine Komponente auf Schicht Lj einen Abwärtsruf (downcall) zu einer Komponente auf einer niedrigeren Ebene Li (mit
i < j) tätigen kann und in der Regel eine Antwort erwartet. Nur in Ausnahmefällen wird ein Aufwärtsruf (upcall) zu einer
höheren Ebene gemacht. Die drei häufigen Fälle sind in Abbildung 2.1 dargestellt.

Abbildung 2.1(a) zeigt eine Standardorganisation, bei der nur Aufrufe zur nächsten niedrigeren Schicht gemacht
werden. Diese Organisation wird häufig bei Netzwerkkommunikation eingesetzt.

In vielen Situationen stossen wir auch auf die in Abbildung 2.1(b) gezeigte Organisation. Betrachten Sie zum Beispiel
eine Anwendung A, die eine Bibliothek LOS verwendet, um eine Schnittstelle zum Betriebssystem herzustellen. Gleichzeitig
nutzt die Anwendung eine spezialisierte mathematische Bibliothek Lmath, die ebenfalls unter Verwendung von LOS
implementiert wurde. In diesem Fall, bezogen auf Abbildung 2.1(b), ist A auf Schicht N − 1, Lmath auf Schicht N − 2 und
LOS, die beiden gemeinsam ist, auf Schicht N − 3 implementiert.

Schliesslich wird in Abbildung 2.1(c) eine besondere Situation gezeigt. In einigen Fällen ist es praktisch, eine untere
Schicht einen Aufruf zu ihrer nächsthöheren Schicht machen zu lassen. Ein typisches Beispiel ist, wenn ein
Betriebssystem das Auftreten eines Ereignisses signalisiert, zu diesem Zweck ruft es eine benutzerdefinierte Operation
auf, für die eine Anwendung zuvor eine Referenz (typischerweise als Handle bezeichnet) übergeben hatte.

#### Geschichtete Kommunikationsprotokolle

Eine bekannte und allgegenwärtig angewandte geschichtete Architektur ist die der Kommunikationsprotokoll-Stacks. Wir
konzentrieren uns hier auf einen Überblick und vertagen eine detaillierte Diskussion auf Abschnitt 4.1.1.

In Kommunikationsprotokoll-Stacks implementiert jede Schicht einen oder mehrere Kommunikationsdienste, die es
ermöglichen, Daten von einer Quelle zu einem oder mehreren Zielen zu senden. Zu diesem Zweck bietet jede Schicht eine
Schnittstelle an, die die aufrufbaren Funktionen spezifiziert. Grundsätzlich sollte die Schnittstelle die tatsächliche
Implementierung eines Dienstes vollständig verbergen. Ein weiteres wichtiges Konzept im Bereich der Kommunikation ist
das eines (Kommunikations-)Protokolls, das die Regeln beschreibt, die die Parteien befolgen werden, um Informationen
auszutauschen. Es ist wichtig, den Unterschied zwischen einem von einer Schicht angebotenen Dienst, der Schnittstelle,
durch die dieser Dienst verfügbar gemacht wird, und dem Protokoll, das eine Schicht implementiert, um die Kommunikation
zu etablieren, zu verstehen. Diese Unterscheidung wird in Abbildung 2.2 gezeigt.

Um diese Unterscheidung klarzumachen, betrachten Sie einen zuverlässigen, verbindungsorientierten Dienst, der von vielen
Kommunikationssystemen bereitgestellt wird. In diesem Fall muss eine kommunizierende Partei zuerst eine Verbindung zu
einer anderen Partei aufbauen, bevor die beiden Nachrichten senden und empfangen können. Zuverlässig bedeutet, dass
starke Garantien gegeben werden, dass gesendete Nachrichten tatsächlich an die andere Seite geliefert werden, selbst
wenn ein hohes Risiko besteht, dass Nachrichten verloren gehen (wie es zum Beispiel bei der Verwendung eines drahtlosen
Mediums der Fall sein kann). Darüber hinaus stellen solche Dienste im Allgemeinen auch sicher, dass Nachrichten in
derselben Reihenfolge geliefert werden, in der sie gesendet wurden.

Diese Art von Dienst wird im Internet durch das Transmission Control Protocol (TCP) realisiert. Das Protokoll
spezifiziert, welche Nachrichten für den Aufbau oder Abbau einer Verbindung ausgetauscht werden müssen, was getan werden
muss, um die Reihenfolge der übertragenen Daten zu bewahren, und was beide Parteien tun müssen, um Daten zu erkennen und
zu korrigieren, die während der Übertragung verloren gegangen sind. Der Dienst wird in Form einer relativ einfachen
Programmierschnittstelle verfügbar gemacht, die Aufrufe zum Aufbau einer Verbindung, zum Senden und Empfangen von
Nachrichten sowie zum erneuten Abbau der Verbindung enthält. Tatsächlich gibt es unterschiedliche Schnittstellen, die
oft vom Betriebssystem oder der verwendeten Programmiersprache abhängen. Ebenso gibt es viele Implementierungen des
Protokolls und seiner Schnittstellen. (Alle Einzelheiten finden Sie in [Stevens, 1994; Wright und Stevens, 1995].)

#### Anwendungsschichtung

Wenden wir uns nun der logischen Schichtung von Anwendungen zu. In Anbetracht der Tatsache, dass eine grosse Klasse
verteilter Anwendungen darauf abzielt, Benutzern oder Anwendungen den Zugang zu Datenbanken zu ermöglichen, haben viele
Menschen eine Unterscheidung zwischen drei logischen Ebenen befürwortet, die im Wesentlichen einem geschichteten
Architekturstil folgen:

- Die Anwendungsschnittstellenebene
- Die Verarbeitungsebene
- Die Datenebene

In Übereinstimmung mit dieser Schichtung sehen wir, dass Anwendungen oft aus ungefähr drei verschiedenen Teilen
konstruiert werden können: einem Teil, der die Interaktion mit einem Benutzer oder einer externen Anwendung handhabt,
einem Teil, der auf einer Datenbank oder einem Dateisystem operiert, und einem mittleren Teil, der im Allgemeinen die
Kernfunktionalität der Anwendung enthält. Dieser mittlere Teil ist logischerweise auf der Verarbeitungsebene platziert.
Im Gegensatz zu Benutzeroberflächen und Datenbanken gibt es auf der Verarbeitungsebene nicht viele gemeinsame Aspekte.
Daher werden wir mehrere Beispiele geben, um diese Ebene zu verdeutlichen.

Als erstes Beispiel betrachten wir eine einfache Internet-Suchmaschine, zum Beispiel eine, die auf den Kauf von Häusern
spezialisiert ist. Solche Suchmaschinen erscheinen als scheinbar einfache Websites, über die jemand Beschreibungen wie
eine Stadt oder Region, eine Preisspanne, den Typ des Hauses usw. angeben kann. Das Backend besteht aus einer riesigen
Datenbank von derzeit zum Verkauf stehenden Häusern. Die Verarbeitungsschicht tut nichts anderes, als die
bereitgestellten Beschreibungen in eine Sammlung von Datenbankabfragen umzuwandeln, die Antworten abzurufen und diese
Antworten nachzubearbeiten, indem sie beispielsweise die Ausgabe nach Relevanz sortiert und anschliessend eine HTML-Seite
generiert. Abbildung 2.4 zeigt diese Organisation.

Als weiteres Beispiel betrachten wir die Organisation der Website dieses Buches, insbesondere die Schnittstelle, die es
jemandem ermöglicht, eine personalisierte digitale Kopie des Buches im PDF-Format zu erhalten. In diesem Fall besteht
die Schnittstelle aus einem WordPress-basierten Webserver, der lediglich die E-Mail-Adresse des Benutzers sammelt (und
einige Informationen darüber, welche Version des Buches angefordert wird). Diese Informationen werden intern an eine
Datei requests.txt angehängt. Die Datenebene ist einfach: Sie besteht lediglich aus einer Sammlung von LaTeX-Dateien und
Abbildungen, die gemeinsam das gesamte Buch ausmachen.

Das Erstellen einer personalisierten Kopie besteht darin, die E-Mail-Adresse des Benutzers in jede der Abbildungen
einzubetten. Zu diesem Zweck wird einmal alle fünf Minuten ein separater Prozess gestartet, der die Liste der Anfragen
durchgeht und nacheinander die E-Mail-Adresse des Anfragenden in jede Bitmap-Grafik einfügt, eine frische Kopie des
Buches generiert, das erzeugte PDF an einem speziellen Ort speichert (zugänglich über eine einzigartige URL) und dem
Anfragenden eine E-Mail sendet, dass die Kopie zum Download bereitsteht. Dieser Prozess wird fortgesetzt, bis alle
Anfragen bearbeitet wurden. In diesem Beispiel sehen wir also, dass die Verarbeitungsschicht die Datenebene oder die
Anwendungsschnittstellenebene in Bezug auf Rechenaufwand und auszuführende Aktionen überwiegt.

### 2.1.2 Serviceorientierte Architekturen

Obwohl der geschichtete Architekturstil beliebt ist, stellt einer seiner grössten Nachteile die oft starke Abhängigkeit
zwischen den verschiedenen Schichten dar. Gute Beispiele, bei denen diese potenziellen Abhängigkeiten sorgfältig
berücksichtigt wurden, findet man bei der Gestaltung von Kommunikationsprotokoll-Stacks. Schlechte Beispiele umfassen
Anwendungen, die im Wesentlichen als Zusammensetzungen bestehender Komponenten entworfen und entwickelt wurden, ohne
viel Rücksicht auf die Stabilität von Schnittstellen oder die Komponenten selbst, ganz zu schweigen von der
Überschneidung der Funktionalität zwischen verschiedenen Komponenten. (Ein überzeugendes Beispiel wird von
Kucharski [2020] gegeben, der die Abhängigkeit von einer einfachen Komponente beschreibt, die eine gegebene Zeichenkette
mit Nullen oder Leerzeichen auffüllt. Der Autor hat die Komponente aus der NPM-Bibliothek zurückgezogen, was Tausende
von Programmen beeinträchtigte.)

Solche direkten Abhängigkeiten von spezifischen Komponenten haben zu einem Architekturstil geführt, der eine lockerere
Organisation in eine Sammlung separater, unabhängiger Einheiten widerspiegelt. Jede Einheit kapselt einen Dienst ein. Ob
sie nun Service, Objects oder Microservices genannt werden, sie haben gemeinsam, dass der Dienst als separater Prozess 
(oder Thread) ausgeführt wird. Natürlich bedeutet das separate Ausführen von Einheiten nicht unbedingt eine Verringerung
der Abhängigkeiten im Vergleich zu einem geschichteten Architekturstil.

#### Objektbasierte Architekturstile

Nehmen wir den objektbasierten Ansatz als Beispiel, haben wir eine logische Organisation, wie in Abbildung 2.5 gezeigt.
Im Wesentlichen entspricht jedes Objekt dem, was wir als Komponente definiert haben, und diese Komponenten sind durch
einen Prozeduraufrufmechanismus verbunden. Im Falle verteilter Systeme kann ein Prozeduraufruf auch über ein Netzwerk
stattfinden, das heisst, das aufrufende Objekt muss nicht auf derselben Maschine ausgeführt werden wie das aufgerufene
Objekt. Tatsächlich kann der genaue Ort, an dem das aufgerufene Objekt sich befindet, für den Anrufer transparent sein:
Das aufgerufene Objekt kann ebenso gut als separater Prozess auf derselben Maschine laufen.

Objektbasierte Architekturen sind attraktiv, weil sie eine natürliche Art bieten, Daten (den sogenannten Zustand eines
Objekts) und die Operationen, die mit diesen Daten durchgeführt werden können (die als Methoden eines Objekts bezeichnet
werden), in eine einzige Einheit einzukapseln. Die von einem Objekt angebotene Schnittstelle verbirgt
Implementierungsdetails, was im Wesentlichen bedeutet, dass wir ein Objekt prinzipiell vollständig unabhängig von seiner
Umgebung betrachten können. Wie bei Komponenten bedeutet dies auch, dass ein Objekt, wenn die Schnittstelle klar
definiert und ansonsten unberührt gelassen wird, durch ein anderes mit derselben Schnittstelle ersetzt werden kann.

Diese Trennung zwischen Schnittstellen und den Objekten, die diese Schnittstellen implementieren, ermöglicht es uns,
eine Schnittstelle auf einer Maschine zu platzieren, während das Objekt selbst auf einer anderen Maschine liegt. Diese
Organisation, die in Abbildung 2.6 dargestellt wird, wird üblicherweise als verteiltes Objekt oder gelegentlich als
entferntes Objekt bezeichnet.

Wenn sich ein Client an ein verteiltes Objekt bindet, wird eine Implementierung der Schnittstelle des Objekts, die als
Proxy bezeichnet wird, in den Adressraum des Clients geladen. Ein Proxy ist analog zu einem sogenannten Client-Stub in
RPC-Systemen. Das Einzige, was es tut, ist, Methodenaufrufe in Nachrichten zu verpacken und Antwortnachrichten
auszupacken, um das Ergebnis des Methodenaufrufs an den Client zurückzugeben. Das eigentliche Objekt befindet sich auf
einem Server, wo es dieselbe Schnittstelle bietet wie auf der Client-Maschine. Eingehende Aufrufanfragen werden zuerst
an einen Server-Stub weitergeleitet, der sie auspackt, um Methodenaufrufe an der Schnittstelle des Objekts auf dem
Server durchzuführen. Der Server-Stumpf ist auch dafür verantwortlich, Rückgabewerte in eine Nachricht zu verpacken und
diese Antwortnachrichten an den Proxy auf der Client-Seite weiterzuleiten.

Das serverseitige Stub wird oft als Skelett bezeichnet, da es die grundlegenden Mittel zur Verfügung stellt, um den
serverseitigen Middleware-Zugriff auf die benutzerdefinierten Objekte zu ermöglichen. In der Praxis enthält es oft
unvollständigen Code in Form einer sprachspezifischen Klasse, die vom Entwickler weiter spezialisiert werden muss.

Ein charakteristisches, aber etwas kontraintuitives Merkmal der meisten verteilten Objekte ist, dass ihr Zustand nicht
verteilt ist: Er befindet sich auf einer einzelnen Maschine. Nur die von dem Objekt implementierten Schnittstellen sind
auf anderen Maschinen verfügbar. Solche Objekte werden auch als entfernte Objekte bezeichnet. In einem allgemeinen
verteilten Objekt kann der Zustand selbst physisch über mehrere Maschinen verteilt sein, aber diese Verteilung ist
ebenfalls vor den Clients hinter den Schnittstellen des Objekts verborgen.

#### Mikroservice-Architekturstil

Man könnte argumentieren, dass objektbasierte Architekturen die Grundlage für die Kapselung von Diensten in unabhängige
Einheiten bilden. Kapselung ist hier das Schlüsselwort: Der Dienst als Ganzes wird als eine in sich geschlossene Einheit
realisiert, obwohl er möglicherweise andere Dienste nutzen kann. Indem verschiedene Dienste klar voneinander getrennt
werden, so dass sie unabhängig voneinander funktionieren können, ebnet man den Weg zu serviceorientierten Architekturen,
die allgemein als SOAs abgekürzt werden.

Angeregt durch objektorientierte Designs und inspiriert durch den Unix-Ansatz, bei dem viele kleine und gegenseitig
unabhängige Programme leicht zusammengesetzt werden können, um grössere Programme zu bilden, haben Softwarearchitekten an
dem gearbeitet, was als Microservices bezeichnet wird. Wesentlich ist, dass jeder Mikroservice als separater (Netzwerk-)
Prozess ausgeführt wird. Die Implementierung eines Mikroservices könnte in Form eines entfernten Objekts erfolgen, aber
das ist keine Voraussetzung. Ausserdem gibt es, obwohl man von Mikroservices spricht, keine allgemeine Übereinstimmung
darüber, wie gross ein solcher Dienst sein sollte. Am wichtigsten ist jedoch, dass ein Mikroservice wirklich einen
separaten, unabhängigen Dienst darstellt. Mit anderen Worten, Modularisierung ist der Schlüssel zur Gestaltung von
Mikroservices [Wolff, 2017].

Dennoch spielt die Grösse eine Rolle. Indem wir bereits festgestellt haben, dass Mikroservices als separate vernetzte
Prozesse laufen, erhalten wir auch eine Wahlmöglichkeit, wo wir einen Mikroservice platzieren. Wie wir später in diesem
Kapitel sehen werden, haben mit dem Aufkommen von Edge- und Fog-Infrastrukturen Diskussionen über die Orchestrierung der
Bereitstellung verteilter Anwendungen über verschiedene Schichten begonnen. Mit anderen Worten, wo platzieren wir was.

#### Grobgliedrige Dienstzusammensetzung

In einer serviceorientierten Architektur wird eine verteilte Anwendung oder ein System im Wesentlichen als
Zusammensetzung vieler Dienste konstruiert. Ein Unterschied (wenn auch nicht strikt) zu Mikroservices besteht darin,
dass nicht alle diese Dienste derselben administrativen Organisation angehören müssen. Wir sind bereits auf dieses
Phänomen gestossen, als wir Cloud-Computing diskutierten: Es kann durchaus sein, dass eine Organisation, die ihre
Geschäftsanwendung betreibt, Speicherdienste eines Cloud-Anbieters nutzt. Diese Speicherdienste sind logisch vollständig
in eine einzelne Einheit eingekapselt, deren Schnittstelle Kunden zur Verfügung gestellt wird.

Natürlich ist Speicherung ein eher grundlegender Dienst, aber es kommen leicht komplexere Situationen in den Sinn.
Betrachten Sie zum Beispiel einen Webshop, der Waren wie E-Books verkauft. Eine einfache Implementierung, die der zuvor
diskutierten Anwendungsschichtung folgt, könnte aus einer Anwendung zur Auftragsbearbeitung bestehen, die wiederum auf
einer lokalen Datenbank mit den E-Books operiert. Die Auftragsbearbeitung umfasst typischerweise das Auswählen von
Artikeln, das Registrieren und Überprüfen des Lieferkanals (eventuell per E-Mail), aber auch die Sicherstellung, dass
eine Zahlung erfolgt. Letzteres kann von einem separaten Dienst, der von einer anderen Organisation betrieben wird,
gehandhabt werden, zu dem ein kaufender Kunde für die Zahlung umgeleitet wird, wonach die E-Book-Organisation
benachrichtigt wird, damit sie die Transaktion abschliessen kann. Dieses Beispiel zeigt auch, dass, während Mikroservices
als relativ klein angesehen werden, ein allgemeiner Dienst als relativ gross erwartet werden kann. Tatsächlich ist es
nicht ungewöhnlich, einen Dienst als Sammlung von Mikroservices zu implementieren.

Auf diese Weise sehen wir, dass das Problem der Entwicklung eines verteilten Systems teilweise eines der
Dienstzusammensetzung ist und sicherzustellen, dass diese Dienste harmonisch zusammenarbeiten. Tatsächlich ist dieses
Problem völlig analog zu den Problemen der Unternehmensanwendungsintegration, die in Abschnitt 1.3.2 diskutiert wurden.
Entscheidend ist und bleibt, dass jeder Dienst eine gut definierte (Programmier-)Schnittstelle bietet. In der Praxis
bedeutet dies auch, dass jeder Dienst seine eigene Schnittstelle anbietet, was die Zusammensetzung von Diensten alles
andere als trivial machen kann.

#### Ressourcenbasierte Architekturen

Als immer mehr Dienste über das Web verfügbar wurden und die Entwicklung verteilter Systeme durch Dienstzusammensetzung
wichtiger wurde, begannen Forscher, die Architektur von meist webbasierten verteilten Systemen neu zu überdenken. Eines
der Probleme bei der Dienstzusammensetzung besteht darin, dass die Verbindung verschiedener Komponenten leicht zu einem
Integrationsalbtraum werden kann.

Als Alternative kann man ein verteiltes System auch als eine riesige Sammlung von Ressourcen betrachten, die einzeln von
Komponenten verwaltet werden. Ressourcen können von (fernen) Anwendungen hinzugefügt oder entfernt und ebenso abgerufen
oder geändert werden. Dieser Ansatz wurde mittlerweile für das Web weitgehend übernommen und ist als Representational
State Transfer (REST) bekannt [Fielding, 2000]. Es gibt vier Schlüsselmerkmale von sogenannten
RESTful-Architekturen [Pautasso et al., 2008]:

1. Ressourcen werden durch ein einziges Namensschema identifiziert
2. Alle Dienste bieten dieselbe Schnittstelle, bestehend aus höchstens vier Operationen, wie in Abbildung 2.7 gezeigt
3. Nachrichten, die an einen Dienst gesendet oder von ihm empfangen werden, sind vollständig selbstbeschreibend
4. Nach der Ausführung einer Operation bei einem Dienst vergisst diese Komponente alles über den Anrufer

Die letzte Eigenschaft wird auch als zustandslose Ausführung bezeichnet, ein Konzept, auf das wir in Kapitel 3
zurückkommen. Um zu veranschaulichen, wie RESTful in der Praxis funktionieren kann, betrachten Sie einen
Cloud-Speicherdienst wie Amazons Simple Storage Service (Amazon S3). Amazon S3, beschrieben in [Murty, 2008] und
neuerdings in [Culkin und Zazon, 2022], unterstützt zwei Ressourcen: Objekte, die im Wesentlichen das Äquivalent zu
Dateien sind, und Buckets, das Äquivalent zu Verzeichnissen. Es gibt kein Konzept, Buckets in Buckets zu platzieren. Ein
Objekt mit dem Namen ObjectName in einem Bucket BucketName wird durch den folgenden Uniform Resource Identifier (URI)
bezeichnet:

`https://s3.amazonaws.com/BucketName/ObjectName`

Um einen Bucket oder ein Objekt zu erstellen, würde eine Anwendung im Wesentlichen eine PUT-Anfrage mit der URI des
Buckets/Objekts senden. Grundsätzlich wird bei dem Dienst das HTTP-Protokoll verwendet. Mit anderen Worten, es handelt
sich einfach um eine weitere HTTP-Anfrage, die anschliessend von S3 korrekt interpretiert wird. Wenn der Bucket oder das
Objekt bereits existiert, wird eine HTTP-Fehlermeldung zurückgegeben.

Ähnlich, um zu erfahren, welche Objekte in einem Bucket enthalten sind, würde eine Anwendung eine GET-Anfrage mit der
URI dieses Buckets senden. S3 gibt dann eine Liste der Objektnamen zurück, wiederum als gewöhnliche HTTP-Antwort.

Die RESTful-Architektur ist aufgrund ihrer Einfachheit beliebt geworden. Pautasso et al. [2008] haben RESTful-Dienste
mit dienstspezifischen Schnittstellen verglichen und, wie zu erwarten, haben beide ihre Vor- und Nachteile. Insbesondere
kann die Einfachheit von RESTful-Architekturen leicht Lösungen für komplexe Kommunikationsschemata verhindern. Ein
Beispiel sind verteilte Transaktionen, die in der Regel erfordern, dass Dienste den Ausführungszustand verfolgen.
Andererseits gibt es viele Beispiele, in denen RESTful-Architekturen perfekt zu einem einfachen Integrationsschema von
Diensten passen, während die Vielzahl von Dienstschnittstellen die Angelegenheit komplizieren kann.