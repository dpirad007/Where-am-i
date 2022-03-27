<script>
    import { Badge, Breadcrumb, Card } from "spaper";

    import "../global.css";

    const items = [
        {
            text: "Where-am-i ?",
            href: "/",
        },
        {
            text: "Profile",
            href: "/profile",
        },
    ];

    let rollNo = 501851;
    let name = "Ved Patil";
    let father = "Arpit Bhat";
    let mother = "Archit Bhonsle";
    let contact = "98989839389";
    let address = "Pen, Pen, Pen";
</script>

<main>
    <Breadcrumb {items} border="True" />
    <h4>
        Profile <Badge type="secondary">{name}</Badge>
        <Badge type="success">Roll {rollNo}</Badge>
    </h4>
    <div class="profile-content">
        <div class="profile-card-center">
            <Card
                class="profile-card"
                title="Ved Patil"
                subTitle={rollNo}
                image="https://picsum.photos/768"
                ><p>Father: {father}</p>
                <p>Mother: {mother}</p>
                <p>Contact: {contact}</p>
                <p>Address: {address}</p></Card
            >
        </div>
    </div>
</main>

<style>
    .profile-content {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 15px;
    }
    .profile-card-center {
        width: 500px;
    }
</style>
