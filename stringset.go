package stringset

type StringSet struct {
	storage map[string]interface{}
}

func New(values ...string) *StringSet {
	storage := make(map[string]interface{}, len(values))
	s := &StringSet{storage}
	for _, v := range values {
		s.storage[v] = struct{}{}
	}
	return s
}

func (s *StringSet) Add(k string) {
	s.storage[k] = struct{}{}
}

func (s *StringSet) Get(k string) (string, bool) {
	v, ok := s.storage[k]
	if !ok {
		return "", false
	}
	return v.(string), ok
}

func (s *StringSet) Delete(k string) {
	delete(s.storage, k)
}

func (s *StringSet) Contains(v string) bool {
	_, ok := s.storage[v]
	return ok
}

func (s *StringSet) ContainsAny(v ...string) bool {
	for _, k := range v {
		if _, ok := s.storage[k]; ok {
			return true
		}
	}
	return false
}

func (s *StringSet) ContainsAll(v ...string) bool {
	for _, k := range v {
		if _, ok := s.storage[k]; !ok {
			return false
		}
	}
	return true
}

func (s *StringSet) Storage() *map[string]interface{} {
	return &s.storage
}

func (s *StringSet) Copy() StringSet {
	ret := make(map[string]interface{}, len(s.storage))
	for k := range s.storage {
		ret[k] = struct{}{}
	}
	return StringSet{ret}
}

func (s *StringSet) Length() int {
	return len(s.storage)
}

func (s *StringSet) Empty() bool {
	return len(s.storage) == 0
}

func (s *StringSet) Values() []string {
	ret := make([]string, len(s.storage))
	i := 0
	for v := range s.storage {
		ret[i] = v
	}
	return ret
}

func (s *StringSet) Union(other *StringSet) StringSet {
	ret := s.Copy()
	for k := range *other.Storage() {
		ret.storage[k] = struct{}{}
	}
	return ret
}

func (s *StringSet) Intersection(other *StringSet) StringSet {
	storage := make(map[string]interface{}, 0)
	for k := range s.storage {
		if other.Contains(k) {
			storage[k] = struct{}{}
		}
	}
	return StringSet{storage}
}

func (s *StringSet) Difference(other *StringSet) StringSet {
	storage := make(map[string]interface{}, 0)
	for k := range s.storage {
		if !other.Contains(k) {
			storage[k] = struct{}{}
		}
	}
	return StringSet{storage}
}
