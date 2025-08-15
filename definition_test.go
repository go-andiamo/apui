package apui

import (
	"github.com/go-andiamo/chioas"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestBrowser_findRequestDef(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		b := &Browser{
			definition: &chioas.Definition{
				Paths: chioas.Paths{
					"/myapi": {
						Tag: "root",
						Paths: chioas.Paths{
							"/pets": {
								Tag: "pets",
							},
						},
					},
				},
			},
		}
		r := httptest.NewRequest("GET", "/myapi/pets", nil)
		path, paths, parts := b.findRequestDef(r)
		require.NotNil(t, path)
		require.Equal(t, "pets", path.Tag)
		require.Len(t, paths, 2)
		require.NotNil(t, paths[0])
		require.Equal(t, "root", paths[0].Tag)
		require.NotNil(t, paths[1])
		require.Equal(t, "pets", paths[1].Tag)
		require.Equal(t, []string{"myapi", "pets"}, parts)
	})
	t.Run("match path param", func(t *testing.T) {
		b := &Browser{
			definition: &chioas.Definition{
				Paths: chioas.Paths{
					"/myapi": {
						Tag: "root",
						Paths: chioas.Paths{
							"/pets": {
								Tag: "pets",
								Paths: chioas.Paths{
									"/{id}": {
										Tag: "pet",
									},
								},
							},
						},
					},
				},
			},
		}
		r := httptest.NewRequest("GET", "/myapi/pets/123", nil)
		path, paths, parts := b.findRequestDef(r)
		require.NotNil(t, path)
		require.Equal(t, "pet", path.Tag)
		require.Len(t, paths, 3)
		require.NotNil(t, paths[0])
		require.Equal(t, "root", paths[0].Tag)
		require.NotNil(t, paths[1])
		require.Equal(t, "pets", paths[1].Tag)
		require.NotNil(t, paths[2])
		require.Equal(t, "pet", paths[2].Tag)
		require.Equal(t, []string{"myapi", "pets", "123"}, parts)
	})
	t.Run("match path param regex", func(t *testing.T) {
		b := &Browser{
			definition: &chioas.Definition{
				Paths: chioas.Paths{
					"/myapi": {
						Tag: "root",
						Paths: chioas.Paths{
							"/pets": {
								Tag: "pets",
								Paths: chioas.Paths{
									"/{id:^\\d{3}$}": {
										Tag: "pet",
									},
								},
							},
						},
					},
				},
			},
		}
		r := httptest.NewRequest("GET", "/myapi/pets/123", nil)
		path, paths, parts := b.findRequestDef(r)
		require.NotNil(t, path)
		require.Equal(t, "pet", path.Tag)
		require.Len(t, paths, 3)
		require.NotNil(t, paths[0])
		require.Equal(t, "root", paths[0].Tag)
		require.NotNil(t, paths[1])
		require.Equal(t, "pets", paths[1].Tag)
		require.NotNil(t, paths[2])
		require.Equal(t, "pet", paths[2].Tag)
		require.Equal(t, []string{"myapi", "pets", "123"}, parts)
	})
	t.Run("match path param with sub-path", func(t *testing.T) {
		b := &Browser{
			definition: &chioas.Definition{
				Paths: chioas.Paths{
					"/myapi": {
						Tag: "root",
						Paths: chioas.Paths{
							"/pets": {
								Tag: "pets",
								Paths: chioas.Paths{
									"/{id}": {
										Tag: "pet",
										Paths: chioas.Paths{
											"/name": {
												Tag: "pet name",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		r := httptest.NewRequest("GET", "/myapi/pets/123/name", nil)
		path, paths, parts := b.findRequestDef(r)
		require.NotNil(t, path)
		require.Equal(t, "pet name", path.Tag)
		require.Len(t, paths, 4)
		require.NotNil(t, paths[0])
		require.Equal(t, "root", paths[0].Tag)
		require.NotNil(t, paths[1])
		require.Equal(t, "pets", paths[1].Tag)
		require.NotNil(t, paths[2])
		require.Equal(t, "pet", paths[2].Tag)
		require.NotNil(t, paths[3])
		require.Equal(t, "pet name", paths[3].Tag)
		require.Equal(t, []string{"myapi", "pets", "123", "name"}, parts)
	})
	t.Run("no definition", func(t *testing.T) {
		b := &Browser{}
		r := httptest.NewRequest("GET", "/example", nil)
		path, paths, parts := b.findRequestDef(r)
		require.Nil(t, path)
		require.Len(t, paths, 1)
		require.Nil(t, paths[0])
		require.Equal(t, []string{"example"}, parts)
	})
	t.Run("no match", func(t *testing.T) {
		b := &Browser{
			definition: &chioas.Definition{},
		}
		r := httptest.NewRequest("GET", "/myapi/pets", nil)
		path, paths, parts := b.findRequestDef(r)
		require.Nil(t, path)
		require.Len(t, paths, 2)
		require.Nil(t, paths[0])
		require.Nil(t, paths[1])
		require.Equal(t, []string{"myapi", "pets"}, parts)
	})
}
