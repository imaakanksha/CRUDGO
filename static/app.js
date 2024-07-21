$(document).ready(function() {
    // Function to load items when the page loads
    function loadItems() {
        $.get("/items", function(data) {
            $("#itemList").empty();
            data.forEach(function(item) {
                $("#itemList").append("<p>" + item.Name + " - $" + item.Price.toFixed(2) + " <button onclick='deleteItem(" + item.ID + ")'>Delete</button></p>");
            });
        });
    }

    // Function to handle form submission for adding new items
    $("#addItemForm").submit(function(event) {
        event.preventDefault();
        var newItem = {
            name: $("#itemName").val(),
            price: parseFloat($("#itemPrice").val())
        };

        $.post("/items", JSON.stringify(newItem), function() {
            loadItems(); // Reload items after adding a new one
            $("#itemName").val(""); // Clear input fields
            $("#itemPrice").val("");
        });

        return false;
    });

    // Function to handle deleting items
    window.deleteItem = function(id) {
        $.ajax({
            url: "/items",
            type: "DELETE",
            contentType: "application/json",
            data: JSON.stringify({id: id}),
            success: function() {
                loadItems(); // Reload items after deleting one
            }
        });
    }


    // Function to handle updating items
    $("#updateItemForm").submit(function(event) {
        event.preventDefault();
        var updatedItem = {
            id: parseInt($("#itemId").val()),
            name: $("#updatedItemName").val(),
            price: parseFloat($("#updatedItemPrice").val())
        };

        $.ajax({
            url: "/items",
            type: "PUT",
            contentType: "application/json",
            data: JSON.stringify(updatedItem),
            success: function() {
                loadItems(); // Reload items after updating
                $("#updateItemForm").hide(); // Hide the update form after updating
                $("#updatedItemName").val(""); // Clear input fields
                $("#updatedItemPrice").val("");
            }
        });

        return false;
    });

    // Function to display update form for a specific item
    window.displayUpdateForm = function(id, name, price) {
        $("#itemId").val(id);
        $("#updatedItemName").val(name);
        $("#updatedItemPrice").val(price);
        $("#updateItemForm").show();
    }


    data.forEach(function(item) {
        $("#itemList").append("<p>" + item.Name + " - $" + item.Price.toFixed(2) + " <button onclick='deleteItem(" + item.ID + ")'>Delete</button> <button onclick='displayUpdateForm(" + item.ID + ", \"" + item.Name + "\", " + item.Price + ")'>Update</button></p>");
    });

    
    
    // Load items when the page loads
    loadItems();
});
