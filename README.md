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

Currently the adapted noder interface only allows to filter for IDs. For generic interfaces, we need to support filtering and ordering on the interfaces fields. Also pagination is important. Following query should be possible and supported by ent:

```
query {
    vehicles(filter: {name: "Audi"}) {
        id
        name
        ... on Car {
            wheelPressure
        }
    }
}

query {
    vehicles(first: 3, orderBy: {direction: DESC, field: NAME}) {
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

type Car struct {
	ent.Schema
}

// ent.Schema needs to have new mathod "Interfaces"
func (c Car) Interfaces() []ent.Interf {
	return []ent.Interf{
		VehicleInterface{},
	}
}
```

**Note**: it should be allowed to add edges from and to interfaces (also interface <-> interface edges) !

```go
func (VehicleInterface) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tank", Tank.Type).Unique(),
	}
}

type Tank struct {
    ent.Schema
}

func (Tank) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("vehicle", VehicleInterface.Type).Ref("tank").Unique().Required()
    }
}

```

### allow additional mutatations for interfaces

- **updateVehicle**: We should be able to filter vehicles by their properties and update their fields (ignoring fields from Car,Plane)
- **deleteVehicle**: We should be able to delete vehicles by filtering for their properties. When a Vehicle is deleted, the Car/Plane should also be deleted. But as its the same row, this shold be trivial.
