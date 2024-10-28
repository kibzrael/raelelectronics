package utils

type Color struct {
	Name string
	Hex  string
}

var DisplayTypes = []string{
	"LED IPS",
	"OLED",
	"LED TN",
}

var DisplaySizes = []uint64{
	12,
	13,
	14,
	15,
	16,
}

var Colors = []Color{
	{Name: "Black", Hex: "#000000"},
	{Name: "Silver", Hex: "#000000"},
	{Name: "White", Hex: "#FFFFFF"},
	{Name: "Gray", Hex: "#000000"},
	{Name: "Gold", Hex: "#000000"},
	{Name: "Blue", Hex: "#000000"},
	{Name: "Orange", Hex: "#000000"},
}

var GPUBrands = []string{
	"Intel",
	"Nvidia",
	"Amd",
	"Apple",
	"Arm",
}

var CPUs = []string{
	"Intel i5",
	"Intel i7",
	"Intel Ultra 9",
	"Amd Ryzen 7",
	"Amd Ryzen 9",
	"Arm-based Apple",
	"Arm-based Snapdragon",
}

var BRANDS []string = []string{"Lenovo", "HP", "Dell", "Apple", "Samsung", "Asus", "MSI", "Acer", "Microsoft", "Alienware", "Huawei", "LG", "Gigabyte", "Razer", "Framework", "Fujitsu", "Xiaomi", "Google"}

var MemorySizes = []uint64{
	8,
	12,
	16,
	32,
}

var StorageSizes = []uint64{
	128,
	256,
	512,
	1024,
}

type BrandFilter struct {
	Logo string
	Name string
}

var BrandFilters = []BrandFilter{
	{Logo: "/static/images/brands/lenovo.svg", Name: "Lenovo"},
	{Logo: "/static/images/brands/apple.svg", Name: "Apple"},
	{Logo: "/static/images/brands/hp.svg", Name: "HP"},
	{Logo: "/static/images/brands/msi.svg", Name: "MSI"},
	{Logo: "/static/images/brands/dell.svg", Name: "Dell"},
	{Logo: "/static/images/brands/asus.svg", Name: "Asus"},
	{Logo: "/static/images/brands/acer.svg", Name: "Acer"},
	{Logo: "/static/images/brands/alienware.svg", Name: "Alienware"},
}

var ColorMaps = map[string]string{
	"Silver":   "#CFCFCF",
	"Black":    "#000000",
	"White":    "#FFFFFF",
	"Gray":     "#717378",
	"Gold":     "#FFBB00",
	"Blue":     "#005AFF",
	"Orange":   "#FF9900",
	"Green":    "#6EC531",
	"Brown":    "#856256",
	"Graphite": "#41424C",
	"Red":      "#E31937",
}
