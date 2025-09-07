mocks:
	@mockgen -destination mocks/http_mock.go -typed -package mocks net/http RoundTripper
	@mockgen -source=internal/service/weatherservice/weatherservice.go -destination=mocks/mock_service.go -package=mocks
	@mockgen -source=internal/infrastructure/cache/weathercache/weathercache.go -destination=mocks/mock_cache.go -package=mocks