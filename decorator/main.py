# Component interface
class Coffee:
	def cost(self):
		pass

# Concrete Component
class SimpleCoffee(Coffee):
	def cost(self):
		return 2

# Decorators
def milk_decorator(coffee):
	def wrapper():
		return coffee.cost() + 1
	return wrapper

def sugar_decorator(coffee):
	def wrapper():
		return coffee.cost() + 0.5
	return wrapper

def syrup_decorator(coffee):
	def wrapper():
		return coffee.cost() + 0.7
	return wrapper

# Usage
@syrup_decorator
@sugar_decorator
def coffee_with_syrup_and_sugar():
	return SimpleCoffee()

@milk_decorator
def coffee_with_milk():
	return SimpleCoffee()

print("Cost of coffee with syrup and sugar:", coffee_with_syrup_and_sugar())
print("Cost of coffee with milk:", coffee_with_milk())
