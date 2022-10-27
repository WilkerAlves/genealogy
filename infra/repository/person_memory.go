package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type PersonRepositoryMemory struct {
	DBMemory map[int]*entity.Person
}

func (p *PersonRepositoryMemory) Add(ctx context.Context, name string) (*entity.Person, error) {
	if name == "ERROR" {
		return nil, fmt.Errorf("error for create person with name: %s", name)
	}

	if p.DBMemory == nil {
		p.DBMemory = make(map[int]*entity.Person, 0)
	}

	nsec := int(time.Now().UnixNano())

	person := &entity.Person{
		ID:   nsec,
		Name: name,
	}
	p.DBMemory[nsec] = person
	return person, nil
}

func (p *PersonRepositoryMemory) Get(ctx context.Context, id int) (*entity.Person, error) {
	person, ok := p.DBMemory[id]
	if !ok {
		return nil, fmt.Errorf("person not found in id: %d", id)
	}

	if person == nil {
		return nil, fmt.Errorf("person not found. id: %d", id)
	}

	return person, nil
}

func (p *PersonRepositoryMemory) Update(ctx context.Context, id int, name string) error {
	person, err := p.Get(ctx, id)
	if err != nil {
		return err
	}

	person.Name = name
	p.DBMemory[person.ID] = person
	return nil
}

func (p *PersonRepositoryMemory) Delete(ctx context.Context, id int) error {
	person, err := p.Get(ctx, id)
	if err != nil {
		return err
	}

	delete(p.DBMemory, person.ID)

	return nil
}
