# Usamos Helm para realizar un deploy de Prometheus y Grafana en nuestro cluster K8S local
# Además levantaremos una instancia de la App de ejemplo en Go
install:
	@helm install --name prometheus --set-file extraScrapeConfigs=prometheus/extraScrapeConfigs.yaml stable/prometheus
	@helm install --name grafana stable/grafana
	@helm install --name app app

# Para publicar el puerto de Prometheus en el Host y poder acceder desde http://localhost:9090
forward-prometheus:
	@kubectl port-forward svc/prometheus-server 9090:80

# Para publicar el puerto de Grafana en el Host y poder acceder desde http://localhost:8080
# Para conseguir las credenciales de usuario `admin` usa `make grafana-get-password`
forward-grafana:
	@kubectl port-forward svc/grafana 8080:80

# Para publicar el puerto de la App en el Host y poder acceder desde http://localhost:8081
# Para publicar el puerto de la App en el Host y poder acceder desde http://localhost:8081
forward-app:
	@kubectl port-forward svc/app 8081:8081

# El password de `admin` se genera de forma aleatoria y se guarda como secret
grafana-get-password:
	@kubectl get secret grafana -o jsonpath="{.data.admin-password}" | base64 --decode; echo

# Para los servicios
stop:
	@helm delete --purge prometheus
	@helm delete --purge grafana
	@helm delete --purge app