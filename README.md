# Sanitize
sanitize is a lightweight Golang package designed to clean, normalize, and transform input payloads using json struct tags. This library focuses not on validation, but on data hygiene—ensuring that incoming payloads are properly trimmed, standardized, and ready for use.

## ✨ Features
- Declarative transformations via struct tags
- Common sanitizers: trimming spaces, lowercasing, uppercasing, removing special characters, etc.
- Extensible design for custom sanitization rules
- Works seamlessly with JSON decoding and Go structs

## 🚀 Use Case
Perfect for APIs and services where clean input matters—from user registration forms to structured payloads—sanitize helps enforce consistency and reduce noisy data before validation or persistence.
