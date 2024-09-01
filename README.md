# In-Memory-Key-Value-Store

## Overview

This project implements an in-memory key-value store with thread-safe operations, designed to manage complex objects where each key maps to an object containing attributes and corresponding values. The store ensures that attribute data types remain consistent once defined, and provides several operations for managing and querying the stored data.

## Features

- **Thread-Safe Operations**: All operations are thread-safe, ensuring data integrity in concurrent environments.
- **Flexible Value Types**: The store supports attributes with string, integer, double, and boolean values.
- **CRUD Operations**: Provides basic operations including `get`, `put`, `delete`, `search`, and `keys`.
- **Data Type Enforcement**: Once an attribute's data type is set, any attempt to change it will result in an exception.

## Data Structure

- **Key**: A string that uniquely identifies each entry in the store.
- **Value**: A map/object that contains attributes and their corresponding values.

### Example Structure

```json
{
  "sde_bootcamp": {
    "title": "Hello World",
    "price": 1500.00,
    "enrolled": false,
    "estimated_time": 90
  }
}
