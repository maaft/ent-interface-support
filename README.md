# ent-interface-support

## Start Server

make run

## Reset database

make reset_db

## Add Plane

!! take note of returned IDs !!

```
mutation AddPlane {
  addPlane(
    input: [{ name: "Boeing", description: "yeah !", altitude: 10000 }]
  ) {
    addPlane {
      id
      createdAt
    }
  }
}
```

## Add Car

```
mutation AddCar {
  addCar(input: [{ name: "Audi", description: "lol !", wheelPressure: 3 }]) {
    addCar {
      id
      createdAt
    }
  }
}
```

## Query Vehicles

```
query GetVehicle($id: ID!) {
  vehicle(id: $id) {
    id
    ... on Car {
      wheelPressure
    }
    ... on Plane {
      altitude
    }
  }
}

query QueryVehicle($ids: [ID!]!) {
  vehicles(ids: $ids) {
    id
    name
    ... on Car {
      wheelPressure
    }
    ... on Plane {
      altitude
    }
  }
}
```

## TODOs

### find better method than pulids to get type from ID

Maybe the type-ID pairs should be stored in an extra table?

### support filtering, ordering and pagination for `vehicles` query

Currently the adapted noder interface only allows to filter for IDs. For generic interfaces, we need to support filtering and ordering on the interfaces fields. Also pagination is important.

### allow vehicles on edges of other types

following GraphQL type should be possible:

```
type Garage {
    name: String!
    vehicles: [Vehicle!]!
}

query {
    getGarage(id: "foobar") {
        name
        vehicles {
            id
            ... on Car {
                wheelPressure
            }
            ... on Plane {
                altitude
            }
        }
    }
}
```

This means, that the ent schema also needs to have some definition of what is an "interface schema-type" and what is a "normal schema-type". See next paragraph.

### find suitable ent implementation for code-gen

Everything in this repo is done manually so far. We should find an implementation that is suitable for ents codegen framework.

Instead of using `mixin.Schema`, we should introduce another type `interf.Schema`:

```go
type VehicleInterface struct {
	interf.Schema
}

func (VehicleInterface) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Optional(),
	}
}

// ent.schema needs to have new mathod "Interfaces"
type Car struct {
	ent.Schema
}

func (c Car) Interfaces() []ent.Interf {
	return []ent.Interf{
		VehicleInterface{},
	}
}
```

### allow additional mutatations for interfaces

- **updateVehicle**: We can filter vehicles by their properties and update their fields
- **deleteVehicle**: We should be able to delete vehicles by filtering for their properties
