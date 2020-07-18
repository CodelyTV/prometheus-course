package prometheusdemo;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;
import io.prometheus.client.Counter;
import io.prometheus.client.Gauge;
import io.prometheus.client.exporter.HTTPServer;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.util.Timer;
import java.util.TimerTask;

public class App {
    static final Counter counter = Counter.build()
                                          .name("test_counter")
                                          .labelNames("code")
                                          .help("Example of a counter.")
                                          .register();
    static final Gauge   gauge   = Gauge.build()
                                        .name("test_gauge")
                                        .help("Example of a gauge")
                                        .labelNames("method", "code")
                                        .register();

    public static void main(String[] args) throws IOException {
        scheduleIncrements();

        HttpServer server = HttpServer.create(new InetSocketAddress(8082), 0);
        server.createContext("/send", new SendMetricsHandler());
        server.setExecutor(null);
        server.start();

        new HTTPServer(8081);
    }

    private static void scheduleIncrements() {
        new Timer().scheduleAtFixedRate(new TimerTask() {
            @Override
            public void run() {
                counter.labels("code").inc(200);
            }
        }, 1000, 1000);

        new Timer().scheduleAtFixedRate(new TimerTask() {
            @Override
            public void run() {
                counter.labels("code").inc(400);
            }
        }, 1000, 4000);
    }

    static class SendMetricsHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange http) throws IOException {
            gauge.labels("get", "200").set(Math.random());
            gauge.labels("post", "300").inc();
            http.sendResponseHeaders(201, 0);
            http.getResponseBody().close();
        }
    }
}
