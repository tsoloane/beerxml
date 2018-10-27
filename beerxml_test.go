package beerxml

import (
	"testing"
)

func TestUnmarshalBeerXml(t *testing.T) {
	t.Log("Unmarshal beerxml recipes.xml")
	rec, err := NewBeerXmlFromFile("testfiles/recipes.xml")
	if err != nil {
		t.Error(err)
		return
	}

	expected := 4
	count := len(rec.Recipes)
	if count != expected {
		t.Errorf("Wrong number of recipes. Expected %d go %d", expected, count)
	}

	recipeName := "Porter"
	recipe := rec.Recipes[2]

	if recipe.Name != recipeName {
		t.Errorf("Expected name %s but got %s",
			recipeName, recipe.Name)
	}
}

func TestUnmarshalBeerSmithExport(t *testing.T) {
	t.Log("Unmarshal beerxml drsmurto.xml")
	rec , err := NewBeerXmlFromFile("testfiles/drsmurto.xml")
	if err != nil {
		t.Error(err)
		return
	}
	//Recipe count is 1
	if len(rec.Recipes) != 1 {
		t.Errorf("Expecting 1 recipe, but got %d", len(rec.Recipes))
	}
	//Recipe Name
	recipeName := "Dr Smurto's Golden Ale"
	if rec.Recipes[0].Name != recipeName {
		t.Errorf("Expected '%s', but got %s",
			recipeName, rec.Recipes[0].Name)
	}
}

func TestRecipesRecipe(t *testing.T) {
	//t.Log("Unmarshal test ... ")
	bxml, err := NewBeerXmlFromFile("testfiles/recipes.xml")
	if err != nil {
		t.Error(err)
		return
	}

	if bxml.Recipes[0].Name != "Burton Ale" {
		t.Error("Recipe.Name")
	}

	if bxml.Recipes[0].Hops[0].Hsi != "35.0" {
		t.Error("Recipe.Hops.Hsi")
	}

	if bxml.Recipes[0].Fermentables[0].Amount != 3.628736 {
		t.Error("Recipe.Fermentables.Amount")
	}

	if bxml.Recipes[0].Mash.MashSteps[0].StepTime != 45 {
		t.Error("Recipes.Mash.MashSteps.StepTime")
	}
}
