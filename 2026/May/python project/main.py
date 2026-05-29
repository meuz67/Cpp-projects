if __name__ == "__main__":
    class Person:
        def __init__(self, age):
            self.age = age
        @property
        def age(self):
            return self.age
        @age.setter
        def age(self, value):
            if value < 0 or value > 100:
                raise ValueError("Age must be between 0 and 100")
            else:
                self.age = value
    person = Person(105)
    print(person.age)