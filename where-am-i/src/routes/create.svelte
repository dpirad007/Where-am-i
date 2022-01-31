<script>
    import { Badge, Breadcrumb, Button, Form, Input, Toast } from "spaper";

    import "../global.css";

    let duration = 2;

    const items = [
        {
            text: "Where-am-i ?",
            href: "/",
        },
        {
            text: "Create",
            href: "/create",
        },
    ];

    function displayToast() {
        Toast.open({
            message: "Created",
            type: "success",
            duration: +duration * 1000,
            position: "top-right",
            dismissible: "True",
        });
    }
</script>

<main>
    <Breadcrumb {items} border="True" />
    <h4>Create lecture <Badge type="secondary">start your day</Badge></h4>
    <Form>
        <Input class="create-form-item" />
        <Button
            on:click={displayToast}
            class="create-form-item"
            type="secondary">Create</Button
        >
    </Form>
</main>

<style>
    :global(.create-form-item) {
        margin: 10px;
    }
</style>
