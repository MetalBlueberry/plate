
from TheDimensionsSchema import TheDimensionsSchema

class TheRootSchema():
    def __init__(
        self,
        checked: bool = false,
        dimensions: TheDimensionsSchema = None,
        id: int = 0,
        name: str = "",
        price: float = 0,
        tags: list = [],
    ):
        self.checked: bool = checked
        self.dimensions: dict = dimensions
        self.id: int = id
        self.name: str = name
        self.price: float = price
        self.tags: list = tags
        
