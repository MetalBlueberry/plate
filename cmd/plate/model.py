class TheRootSchema():
    def __init__(
        self,
        checked: bool = false,
        dimensions: TheDimensionsSchema = None,
    ):
        self.checked: bool = checked
        self.dimensions: dict = dimensions


class TheDimensionsSchema():
    def __init__(
        self,
        height: int = 0,
        width: int = 0,
    ):
        self.height: int = height
        self.width: int = width
