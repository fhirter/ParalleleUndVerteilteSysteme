# Verteilte Systeme
Die folgenden Texte sind Kapitel aus dem Buch "M. van Steen and A.S. Tanenbaum, Distributed Systems, 4th ed., distributed-systems.net, 2023".
Eine digitale Kopie kann gratis auf der Website des Buches heruntergeladen werden: [distributed-systems.net](https://www.distributed-systems.net/index.php/books/ds4/). 
Auf Deutsch ist nur die zweite Ausgabe verfügbar, welche aufgrund der Schnelllebigkeit der Materie nicht empfohlen wird.

Die folgenden Texte wurden mit Chat GPT-4 übersetzt und auf vom Dozenten auf Fehler überprüft.

# 1.4 Fallstricke
Es sollte mittlerweile klar sein, dass die Entwicklung eines verteilten Systems eine gewaltige Aufgabe ist.
Wie wir in diesem Buch mehrmals sehen werden, gibt es so viele Probleme zu berücksichtigen, dass das Ergebnis nur Komplexität zu sein scheint.
Dennoch können durch die Befolgung mehrerer Entwurfsprinzipien verteilte Systeme entwickelt werden, die den in diesem Kapitel festgelegten Zielen gut entsprechen. 
Verteilte Systeme unterscheiden sich von traditioneller Software, da Komponenten über ein Netzwerk verteilt sind. 
Wenn diese Dispersion während der Entwurfszeit nicht berücksichtigt wird, werden viele Systeme unnötig komplex und werden Fehler gemacht, die später behoben werden müssen. 
Peter Deutsch, damals bei Sun Microsystems tätig, formulierte folgende falschen Annahmen, die viele bei der Entwicklung einer verteilten Anwendung zum ersten Mal treffen:
- Das Netzwerk ist zuverlässig
- Das Netzwerk ist sicher
- Das Netzwerk ist homogen
- Die Topologie ändert sich nicht
- Latenz ist null
- Bandbreite ist unendlich
- Transportkosten sind null
- Es gibt einen Administrator

Beachten Sie, wie diese Annahmen sich auf Eigenschaften beziehen, die einzigartig für verteilte Systeme sind: Zuverlässigkeit, Sicherheit, Heterogenität und Netzwerk-Topologie; Latenz und Bandbreite; Transportkosten; und schließlich administrative Bereiche.
Bei der Entwicklung von nicht verteilten Anwendungen werden die meisten dieser Probleme wahrscheinlich nicht auftreten. 

Die meisten in diesem Buch besprochenen Prinzipien beziehen sich direkt auf diese Annahmen. 
In allen Fällen diskutieren wir Lösungen für Probleme, die durch die Tatsache verursacht werden, dass eine oder mehrere Annahmen falsch sind. 
Zum Beispiel gibt es einfach keine zuverlässigen Netzwerke, was zur Unmöglichkeit der Erreichung von Ausfallstransparenz führt. 
Wir widmen ein ganzes Kapitel der Tatsache, dass netzwerkgestützte Kommunikation von Natur aus unsicher ist. 
Wir haben bereits argumentiert, dass verteilte Systeme offen sein müssen und die Heterogenität berücksichtigen sollten. 
Ebenso behandeln wir beim Diskutieren von Replikation zur Lösung von Skalierbarkeitsproblemen im Wesentlichen Latenz- und Bandbreitenprobleme. 
Wir werden auch an verschiedenen Stellen in diesem Buch auf Managementprobleme eingehen.

# 1.5 Zusammenfassung
Ein verteiltes System ist eine Sammlung von vernetzten Computersystemen, in denen Prozesse und Ressourcen auf verschiedene Computer verteilt sind. Wir unterscheiden zwischen ausreichend und notwendigerweise verteilt, wobei letzteres sich auf dezentrale Systeme bezieht. Diese Unterscheidung ist wichtig, da das Verteilen von Prozessen und Ressourcen nicht als eigenständiges Ziel betrachtet werden kann. Stattdessen resultieren die meisten Entscheidungen für ein verteiltes System aus dem Bedarf, die Leistung eines einzelnen Computersystems hinsichtlich Zuverlässigkeit, Skalierbarkeit und Effizienz zu verbessern. Aber wenn man bedenkt, dass die meisten zentralisierten Systeme immer noch viel einfacher zu verwalten und zu warten sind, sollte man sich überlegen, Prozesse und Ressourcen zu verteilen. Es gibt auch Fälle, in denen es einfach keine Wahl gibt, zum Beispiel beim Verbinden von Systemen verschiedener Organisationen oder wenn Computer einfach aus verschiedenen Standorten arbeiten (wie beim mobilen Computing).
Entwurfsziele für verteilte Systeme beinhalten das Teilen von Ressourcen und das Sicherstellen von Offenheit. Immer wichtiger wird das Entwerfen von sicheren verteilten Systemen. Darüber hinaus zielen Designer darauf ab, viele der Feinheiten im Zusammenhang mit der Verteilung von Prozessen, Daten und Kontrolle zu verbergen. Diese Verteilungstransparenz kommt jedoch nicht nur mit einem Leistungspreis, in der Praxis kann sie nie vollständig erreicht werden. Die Tatsache, dass Abwägungen zwischen verschiedenen Formen der Verteilungstransparenz getroffen werden müssen, ist inhärent im Design von verteilten Systemen und kann deren Verständnis leicht verkomplizieren. Ein spezifisches, schwieriges Entwurfsziel, das nicht immer gut mit der Erreichung von Verteilungstransparenz harmoniert, ist die Skalierbarkeit. Dies gilt insbesondere für geografische Skalierbarkeit, bei der das Verbergen von Latenzen und Bandbreitenbeschränkungen schwierig sein kann. Ebenso kann die administrative Skalierbarkeit, bei der ein System so gestaltet ist, dass es mehrere administrative Bereiche umfasst, leicht mit Zielen für die Erreichung von Verteilungstransparenz in Konflikt geraten.
Es gibt verschiedene Typen von verteilten Systemen, die als auf die Unterstützung von Berechnungen, Informationsverarbeitung und Durchdringung ausgerichtet klassifiziert werden können. Verteilte Computing-Systeme werden typischerweise für Hochleistungsanwendungen eingesetzt, die oft aus dem Bereich des parallelen Computings stammen. Ein Bereich, der aus der parallelen Verarbeitung entstand, war anfangs das Grid-Computing mit einem starken Fokus auf das weltweite Teilen von Ressourcen, was schließlich zu dem führte, was heute als Cloud-Computing bekannt ist. Cloud-Computing geht über das Hochleistungscomputing hinaus und unterstützt auch verteilte Systeme in traditionellen Büroumgebungen, in denen Datenbanken eine wichtige Rolle spielen. Typischerweise werden in diesen Umgebungen Transaktionsverarbeitungssysteme eingesetzt. Schließlich ist eine aufkommende Klasse von verteilten Systemen dort, wo die Komponenten klein sind, das System ad hoc zusammengestellt wird, vor allem aber nicht mehr durch einen Systemadministrator verwaltet wird. Diese letzte Klasse wird typischerweise von durchdringenden Computingsystemen, einschließlich mobiler Computing-Systeme sowie sensorreichen Umgebungen, dargestellt.
Die Angelegenheiten werden weiter kompliziert durch die Tatsache, dass viele Entwickler ursprünglich Annahmen über das zugrundeliegende Netzwerk treffen, die grundlegend falsch sind. Später, wenn die Annahmen fallen gelassen werden, kann es schwierig sein, unerwünschtes Verhalten zu maskieren. Ein typisches Beispiel ist die Annahme, dass Netzwerk-Latenz nicht signifikant ist. Andere Fallstricke sind die Annahme, dass das Netzwerk zuverlässig, statisch, sicher und homogen ist.