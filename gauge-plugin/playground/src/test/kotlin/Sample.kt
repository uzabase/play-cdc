import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSuite

class Sample {
    private val wireMock = WireMock(8080)

    @BeforeSuite
    fun loadWiremockMappings() {
        wireMock.loadMappingsFrom("fixtures")
    }

    @BeforeScenario
    fun resetRequests() {
        wireMock.resetRequests()
    }
}
